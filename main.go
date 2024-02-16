package main

import (
        "bufio"
        "flag"
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "os/exec"
        "regexp"
        "strings"
)

func searchDomains(domain string, outputFileName string) error {
        resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))
        if err != nil {
                return fmt.Errorf("failed to retrieve data from crt.sh for %s: %v", domain, err)
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return fmt.Errorf("failed to read response body for %s: %v", domain, err)
        }

        re := regexp.MustCompile(`"common_name":"(.*?)"`)
        matches := re.FindAllStringSubmatch(string(body), -1)

        var domainList []string
        for _, match := range matches {
                if strings.Contains(match[1], domain) && !strings.HasPrefix(match[1], "*") {
                        domainList = append(domainList, match[1])
                }
        }

        if len(domainList) == 0 {
                return nil
        }

        cmd := exec.Command("sort", "-u")
        cmd.Stdin = strings.NewReader(strings.Join(domainList, "\n"))
        out, err := cmd.Output()
        if err != nil {
                return fmt.Errorf("failed to run sort -u: %v", err)
        }

        if outputFileName != "" {
                err := ioutil.WriteFile(outputFileName, out, 0644)
                if err != nil {
                        return fmt.Errorf("failed to write to output file %s: %v", outputFileName, err)
                }
                fmt.Printf("Matching domains saved to %s\n", outputFileName)
        } else {
                fmt.Print(strings.TrimSuffix(string(out), "\n")) // Remove newline character
        }

        return nil
}

func main() {
        singleDomain := flag.String("d", "", "Single domain to search")
        inputFile := flag.String("i", "", "Input file containing domains")
        flag.Parse()

        if *singleDomain != "" {
                if err := searchDomains(*singleDomain, ""); err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
                os.Exit(0)
        }

        if *inputFile != "" {
                file, err := os.Open(*inputFile)
                if err != nil {
                        fmt.Printf("Input file '%s' not found: %v\n", *inputFile, err)
                        os.Exit(1)
                }
                defer file.Close()

                scanner := bufio.NewScanner(file)
                for scanner.Scan() {
                        domain := scanner.Text()
                        if err := searchDomains(domain, ""); err != nil {
                                fmt.Println(err)
                                os.Exit(1)
                        }
                }

                if err := scanner.Err(); err != nil {
                        fmt.Printf("Error reading input file: %v\n", err)
                        os.Exit(1)
                }

                os.Exit(0)
        }

        fmt.Println("No options provided. Usage: ./program [-d <domain>] [-i <input_file>]")
        os.Exit(1)
}
