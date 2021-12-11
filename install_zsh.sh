#!/bin/zsh

cd ./src/
go build -o aws-set-profiler
chmod u+x aws-set-profiler

mkdir -p $HOME/.bin
mv aws-set-profiler $HOME/.bin/

cat >> $HOME/.zprofile << EOF

# AWS set profiler
export PATH=\$HOME/.bin:\$PATH
alias aws='aws --profile=\$(cat \$HOME/.aws/profile)'
EOF

cat << EOF
Well done!
Now we put a command \`aws-set-profile\` on \$HOME/.bin.
If you don't have a path to \$HOME/.bin, please add.
You must restart your terminal or call \`source \$HOME/.zshrc\`.
EOF
