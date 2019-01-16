# remote-ssh-commander
Remote SSH Commander = run commands on remote devices via SSH. Specially useful for network engineers

## Settings

Application does not have to much setting. Just set up your commands in command.json file

~~~
[
    {
        "name":"<device_name>",
        "ip": "<device_ip>",
        "username":"<username>",
        "password":"<pasword>",
        "commands":[
            "<command_1>",
            "<command_2>",
            "<command_3>",
            "<command_4>",
            ...
            ""
        ]
    },
    ...
]
~~~

You can add any number of connection and commands. Application will run all of them by line order.

## Running and Parameters

to see all parametes run help

~~~
$ ./rsshc -h

Remote SSH Commander = run commands on remote devices via SSH
usage: rsshc [options]
options:

  -c string
    	commands file path in json format (default "commands.json")
  -h	display this help dialog
  -v	output version information and exit.
~~~

running with commands file from different location

~~~
$ ./rsshc -c ~/Documents/my-commands.json
...
~~~
