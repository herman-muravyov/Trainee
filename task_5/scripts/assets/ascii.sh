#! /bin/bash

source ./scripts/assets/set_color.sh

function display_art()
{
  cat << "EOF"
           
@@@@@@@   @@@@@@  @@@@@@@  @@@@@@  
@@!  @@@ @@!  @@@   @!!   @@!  @@@ 
@!@  !@! @!@!@!@!   @!!   @!@!@!@! 
!!:  !!! !!:  !!!   !!:   !!:  !!! 
:: :  :   :   : :    :     :   : : 
                                   
                                 
 @@@@@@ @@@@@@@@ @@@@@@@  @@@@@@ 
!@@     @@!        @!!   !@@     
 !@@!!  @!!!:!     @!!    !@@!!  
    !:! !!:        !!:       !:! 
::.: :  : :: ::     :    ::.: :  
                                       
EOF
	echo -e "\nHello and welcome."
	echo -e "@Let's calculate the shape."
}