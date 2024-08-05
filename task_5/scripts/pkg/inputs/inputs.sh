#! /bin/bash


function menu_input()
{
	read -p "Enter your choice: " choice
}

function url_input()
{
	read -p "Enter your URL address: " url
}

function output_input()
{
	read -p "Enter output file name: " output
}

function file_extension_input()
{
	read -p "Enter file's extension (ex.: .csv): " ext
}

function continue_using()
{
	echo -e "\n$(set_color "green")Press any button to continue...$(set_color "*")"
	read -r
}