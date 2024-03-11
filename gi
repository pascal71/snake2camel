set -o vi
git init
git add .
git commit -m "Initial commit of snake2camel package"
git tag v0.1.0
git remote add origin git@github.com:pascal71/snake2camel.git
git push -u origin master
git push origin v0.1.0
