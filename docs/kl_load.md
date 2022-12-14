## kl load

load environment variables and mount config files according to defined in kl-config file

### Synopsis

```
Load Environment
This command help you to load environments of the server according to you defined in your kl-config file.

Examples:
  # load environments and mount the configs
  kl load

	# load environments and execute a program with that loaded environments
	kl load <your_cmd>

	# example with npm start
	kl load npm start

	# get environments in json format
	kl load -o json

	# get environments in yaml format
	kl load -o yaml

	# start a new shell with loaded environments
	kl load shell

	# example of env with zsh shell
	kl load zsh
	
```

### Options

```
  -h, --help   help for load
```

### SEE ALSO

* [kl](kl.md)  - kl is command line interface to interact with kloudlite environments

###### Auto generated by kl CLI on 19-October-2022
