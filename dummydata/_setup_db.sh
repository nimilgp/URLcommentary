for FILE in `ls setupfiles`
do
    sudo -u postgres psql -d greenlight -f ./setupfiles/$FILE
    sleep 2
done