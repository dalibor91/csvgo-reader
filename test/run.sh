#!/bin/bash 

location="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
test_data="${location}/data/sample.csv"
test_bin="${location}/bin/test_bin"
test_go="${location}/tests"

function dbg {
	local _run="go build -o ${test_bin} $1"
	echo "$_run"
	export test_data="$test_data" 
	eval "$_run"
	if ! [ $? -eq 0 ];
	then 
		echo "Unable to build"
		exit 1
	fi

	local output="${test_bin}.`date +"%s"`.log"
	eval $test_bin >> $output

	if [ "$DEBUG" = "1" ];
	then 
		cat $output
	fi;

	if ! [ $? -eq 0 ];
	then 
		echo "Unable to execute "
		exit 2
	fi
}

cd $location

for file in $(find $test_go -name '*.go');
do
	echo "================ `date +"%s"` ================="
	dbg $file
done

echo "Tests DONE!"

