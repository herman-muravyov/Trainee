#! /bin/bash

source ./scripts/assets/set_color.sh

function display_art()
{
  cat << "EOF"
  ______   __    __   ______   _______   ________ 
 /      \ /  |  /  | /      \ /       \ /        |
/$$$$$$  |$$ |  $$ |/$$$$$$  |$$$$$$$  |$$$$$$$$/ 
$$ \__$$/ $$ |__$$ |$$ |__$$ |$$ |__$$ |$$ |__    
$$      \ $$    $$ |$$    $$ |$$    $$/ $$    |   
 $$$$$$  |$$$$$$$$ |$$$$$$$$ |$$$$$$$/  $$$$$/    
/  \__$$ |$$ |  $$ |$$ |  $$ |$$ |      $$ |_____ 
$$    $$/ $$ |  $$ |$$ |  $$ |$$ |      $$       |
 $$$$$$/  $$/   $$/ $$/   $$/ $$/       $$$$$$$$/ 
                                                  
EOF
	echo -e "\nHello and welcome."
	echo -e "@Let's calculate the shape."
}