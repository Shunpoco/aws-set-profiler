# aws-set-profiler
This program sets AWS profile which user selected.
You no longer need to call `aws --profile=xxxx`, just call `aws`!

<img src="./img/demo.gif" width="650">

## Install
### Bash
```bash
./install_bash.sh
```

### Zsh
```zsh
./install_zsh.sh
```


Then you restart a terminal.

## How to use
```bash
aws-set-profile
```

Then you select a proper AWS profile.

```bash
$ aws-set-profiler 
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select AWS Profile: 
  ▸ default
    profile2
    profile3
    profile4
```

Now the profile is set!
```bash
$ aws configure list
      Name                    Value             Type    Location
      ----                    -----             ----    --------
   profile                 profile3           manual    --profile
```
