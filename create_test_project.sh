#!/bin/sh

TEMP_DIR=$(mktemp -d)
# Set an environment variable with the name of the temporary directory
export TEST_DIRECTORY="$TEMP_DIR"
cat > "${TEST_DIRECTORY}/configuration.yaml" <<EOF
categories:
	- name: holiday
	  free: yes
	  extraPaid: yes
	- name: out of office
	  free: yes
	  extraPaid: yes
	- name: normal work
	  free: no
	  extraPaid: no
	- name: overtime work
	  free: no
	  extraPaid: yes
EOF
mkdir -p "${TEST_DIRECTORY}/2024/05"
cat >  "${TEST_DIRECTORY}/2024/05/01.timesheet" <<EOF
holiday
EOF
cat >  "${TEST_DIRECTORY}/2024/05/02.timesheet" <<EOF
out of office
EOF
cat >  "${TEST_DIRECTORY}/2024/05/03.timesheet" <<EOF
EOF

echo "$TEST_DIRECTORY"
