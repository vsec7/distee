package main

import (
	"gopkg.in/yaml.v2"
	"bufio"
	"os/user"
	"fmt"
	"io/ioutil"
	"os"
	"flag"
	"strings"
	"github.com/logrusorgru/aurora"
	"github.com/bwmarrin/discordgo"
)

var (
	Setup			bool
	Code			bool
	tok 			string
	chid			string
	path			string
	cfg				string
	Channel_id		string
	msgs 			string
	ConfigFile		string
	Title			string
	config Config
)

type Config struct {
    Token		string	`yaml:"token"`
    DefChan		string	`yaml:"default_channel_id"`
}

func init() {
	flag.BoolVar(&Setup, "setup", false, "Setup Configuration")
	flag.StringVar(&Channel_id, "c", "", "Channel ID")
	flag.StringVar(&Title, "t", "", "Title")
	flag.BoolVar(&Code, "code", false, "Code")
	flag.StringVar(&ConfigFile, "config", "", "Configuration File")
	flag.Usage = func() {
		h := []string{
			"",
			"DisTee (Discord Tee)",
			"",
			"Is a GO tool that works like tee command",
			"Feed input to Discord through the stdin",
			"",
			"Crafted By : github.com/vsec7",
			"",
			"Basic Usage :",
			" ▶ echo \"Hello Cath\" | distee",
			" ▶ cat file.txt | distee -c <channel_id>",
			" ▶ cat file.txt | distee -c <channel_id> -t <title> -code",
			"",
			"Options :",
			"  -setup, --setup                  Setup Configuration",
			"  -c, --c <channel_id>             Send message to custom channel_id",
			"  -t, --t <title>                  Send message with title",
			"  -code, --code                    Send message with code markdown",
			"  -config, --config <config.yaml>  Set custom config.yaml location",
			"",
			"",
		}
		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
	flag.Parse()
}

func Chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println( err )
	}
	    	
    path := usr.HomeDir+"/.distee"
    if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	
	if Setup == true {
    	fmt.Println("----------[ Setup Configuration ]----------")
    	fmt.Println("[?] Enter Discord BOT Token: ")
    	fmt.Scanln(&tok)
    	fmt.Println("[?] Enter Default Channel ID : *Optional")
    	fmt.Scanln(&chid)
    	
    	file, err := os.OpenFile(path+"/config.yaml", os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Failed Creating File: %s", err)
			os.Exit(0)
		}
		buf := bufio.NewWriter(file)
		buf.WriteString("token: "+tok+"\ndefault_channel_id: "+chid+"\n" )
		buf.Flush()
		file.Close()
    	fmt.Println("----------[ Configuration Saved ]----------")
    	os.Exit(0)
    }
	
	if len(ConfigFile) != 0 {
		cfg = ConfigFile
	}else{
		cfg = path+"/config.yaml"
	}
    
    yamlFile, err := ioutil.ReadFile(cfg)
    if err != nil {
		fmt.Printf("[%s] File config.yaml not found!\n", aurora.Red("ERROR"))
        os.Exit(0)
    }
    
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        fmt.Printf("[%s] Cannot parsing config.yaml!\n", aurora.Red("ERROR"))
        os.Exit(0)
    }
    
    if len(Channel_id) != 0 {
    	Channel_id = Channel_id
    } else {
    	Channel_id = config.DefChan
    }

	dg, _ := discordgo.New("Bot "+ config.Token )
	fi, _ := os.Stdin.Stat()

    if (fi.Mode() & os.ModeCharDevice) == 0 {
    	bytes, _ := ioutil.ReadAll(os.Stdin)
        str := string(bytes)
        
	    cs := Chunks(str, 1900)
		for _, c := range cs {
			
			if Code == true {
			    msgs = "```\n" + c + "\n```"
			} else {
				 msgs = "" + c + ""
			}
	
	        if len(Title) != 0 {
		    	msgs = "> " + Title + "\n" + msgs + ""
		    } else {
		    	msgs = "" + msgs + ""
		    }		    
			_, err := dg.ChannelMessageSend(Channel_id, msgs)
			if err != nil {
		        fmt.Printf("[%s] Failed to send message !\n%s\n", aurora.Red("ERROR"), err) 
		        os.Exit(0)
		    }
		}
		
    	fmt.Println(str)
	}
}
