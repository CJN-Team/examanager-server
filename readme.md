# How To Clone With SubModules

git clone --recurse-submodules https://github.com/CJN-Team/Examanager.git

# How To Pull With SubModules

git pull
git submodule update

# How To Add With SubModules

git submodule foreach git add .

# How To Commit (Try one by one) With SubModules

git submodule foreach git commit -m ""

# How To Push With SubModules

git submodule foreach git push

# Nota

Ejecuta estos comandos antes de ejecutar el principal