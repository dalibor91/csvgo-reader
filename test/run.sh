#!/bin/bash 

location="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
test_data="${location}/data/sample.csv"
test_bin="${location}/bin/test_bin"
test_go="${location}/tests"

function dbg {
	local _test_bin="${test_bin}_`date +"%s"`"
	
	local _run="go build -o ${_test_bin} $1"


	echo "$_run"
	export test_data="$test_data" 
	eval "$_run"
	if ! [ $? -eq 0 ];
	then 
		echo "Unable to build"
		exit 1
	fi

	local output="${_test_bin}.`date +"%s"`.log"
	eval $_test_bin >> $output

	if [ "$DEBUG" = "1" ];
	then 
		cat $output
	fi;

	if ! [ $? -eq 0 ];
	then 
		echo "Unable to execute "
		exit 2
	fi

	rm $_test_bin
	rm $output
}

cd $location

files=$(find $test_go -name '*.go')

if [ "$docker" = "1" ];
then 
	echo "Run inside docker with remote library..."
	go get "github.com/dalibor91/csvgo-reader"
	cp -rp "${location}/tests" "${location}/docker"
	files="$(find "${location}/docker" -name '*.go')"

	for file in $files; 
	do 
		sed -i 's/"..\/..\/csv-reader"/"github.com\/dalibor91\/csvgo-reader\/csv-reader"/' $file
	done 
fi

for file in $files;
do
	echo "================ `date +"%s"` ================="
	dbg $file
	echo -e "\n\n"
	sleep 1
done

echo "Tests DONE!"

