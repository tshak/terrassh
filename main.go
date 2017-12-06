package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// SSHKeyPathSuffix is the suffix for the SSH key path
const SSHKeyPathSuffix = "_ssh_key_path"

// SSHUsername is the suffic for the SSH username
const SSHUsername = "_ssh_username"

// HostsSuffix is the suffix for the ip address(s)
const HostsSuffix = "_ssh_hosts"

type sshArgs struct {
	host       string
	sshKeyPath string
	username   string
}

type stringJSONValue struct {
	Value string `json:"value"`
}

type stringArrayJSONValue struct {
	Value []string `json:"value"`
}

func main() {
	validateArgs()
	validateTerraform()

	prefix := strings.TrimSpace(os.Args[1])
	hosts := getHosts(prefix)
	hostIndex := getHostIndex(len(hosts), os.Args)

	sshArgs := sshArgs{
		sshKeyPath: getSSHKeyPath(prefix),
		username:   getSSHUsername(prefix),
		host:       hosts[hostIndex],
	}

	execSSH(sshArgs)
}

func printHelp() {
	fmt.Println(`
USAGE: terrassh <prefix> (hostIndex)

DETAILS:
prefix      Defines the output variable prefix (e.g. "foo" looks for "foo_ssh_hosts" in the terraform output)
hostIndex   The zero based index of "<prefix>_ssh_hosts" to use as the SSH host (Default 0).
`)

	os.Exit(2)
}

func validateArgs() {
	if len(os.Args) == 1 {
		printHelp()
	}
}

func getHostIndex(hostCount int, osArgs []string) int {
	if len(osArgs) >= 3 {
		if hostIndex, err := strconv.Atoi(osArgs[2]); err != nil {
			fatal("hostIndex must be a number")
		} else if hostIndex >= hostCount {
			fatal(fmt.Sprintf("Invalid hostIndex. Only %d hosts specified. Remember that this is a zero based index.", hostCount))
		} else {
			return hostIndex
		}
	}

	return 0
}

func execSSH(args sshArgs) {
	fmt.Printf("Connecting to %s...\n", args.host)
	cmd := exec.Command("ssh", "-i", args.sshKeyPath, fmt.Sprintf("%s@%s", args.username, args.host))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func getSSHKeyPath(prefix string) string {
	return unmarshalValueString(execTerraformOutput(prefix, SSHKeyPathSuffix))
}

func getSSHUsername(prefix string) string {
	return unmarshalValueString(execTerraformOutput(prefix, SSHUsername))
}

func unmarshalValueString(outputContent []byte, outputVar string) string {
	var parsed stringJSONValue
	if err := json.Unmarshal(outputContent, &parsed); err != nil {
		fatal(fmt.Sprintf("Unable to parse %s: %s\nOutput from terraform:\n%s",
			outputVar, err, outputContent))
	}
	return parsed.Value
}

func getHosts(prefix string) []string {
	outputContent, outputVar := execTerraformOutput(prefix, HostsSuffix)

	var parsed stringArrayJSONValue
	if err := json.Unmarshal(outputContent, &parsed); err != nil {
		fatal(fmt.Sprintf("Unable to parse %s: %s\nOutput from terraform:\n%s",
			outputVar, err, outputContent))
	}

	if len(parsed.Value) == 0 {
		fatal(fmt.Sprintf("%s must contain at least one hostname or IP address", outputVar))
	}
	return parsed.Value
}

func validateTerraform() {
	cmd := exec.Command("which", "terraform")
	err := cmd.Run()
	if err != nil {
		fatal("terraform executable not found. Please be sure that it's in your path!")
	}
}

func execTerraformOutput(prefix, suffix string) ([]byte, string) {
	varname := fmt.Sprintf("%s%s", prefix, suffix)

	output, err := exec.Command("terraform", "output", "-json", varname).CombinedOutput()
	if err != nil {
		fatalTerraformOutput(varname, err)
	}

	return output, varname
}

func fatalTerraformOutput(outputVar string, err error) {
	fatal(fmt.Sprintf("Error executing \"terraform output %s\". Ensure that the prefix variables are configured properly. \n%s", outputVar, err))
}

func fatal(message string) {
	fmt.Printf("ERROR: %s\n", message)
	os.Exit(1)
}
