package components

import (
  "fmt"
)

//remeber to create an encode struct
//99,999,999,999

func Documentation(args ...interface{}){
  if len(args) <= 0{
    Help()
  }
  if args == "Errors"{
    ErrorStractures()
  }
}

func ErrorStractures(){
  nb := `
  *****************************************
  **** SOME ERRORS MIGHT BE REDUNDANT   ***
  *****************************************
  `
  err := `
    ERROR while listing agents >ELA
    Error scannig agent rows  >ESAR
    ERROR viewing agent with id of %s ERROR: %s >ESA
    EPIA  Error preparing to insert agent: %s
    ESG Error scanning vovernor with that id
    ERROR viewing president with id of %s ERROR: %s EVP
    ELP Error listing all presidents
    ESPR Error scannig president rows
    ESGR Error scannig governor rows
    EVPSA ERROR viewing votes for polling station %s of agent with id of %s ERROR: %s
    EGCV Error getting candidate votes
    ESCVR Error scannig candidates votes rows: %s
    EPCN Error preparing create news
    ELN Error Listing all news
    ESAN Error sa=scanning all news
    ELHN Error listing handled news
    ELUHN Error listing unhandled news
    ESAUHN Error Scanning for all unhandled news
    ESAHN Error scanning all handled news
    ESAV Error showing all votes
    ESCAV Error scanning for all votes
    EVPV Error viewing polling station votes of agent id %s
    EQWV Error Querying ward votes
    ESWR Error scanning for ward rows
    ELCV Error listing constituency votes
    ESCV Error scannig constituency votes rows: %s
    ESCNV Error scannig county votes rows: %s
    ELNCV Erro listinng county votes
    EZRAWCN Error zero rows affected while creating news
    CAEWIP Creat Agent Error while inserting into president
    CAEGPH ,   ,  , ,   , ,  generating password hash
    CAEIU   ,  ,  , ,   ,  , Inserting into users
    CAECE   ,   , , ,   , , Commit Error
    EVPA ErrorViewing PA
    EPLAPAs Error preparing to list all PAs
    ESAPAs Error scannig All PAs
    EPLPABC Error preparing to PA by creator
    ESPPA Error scanning presidential PAs
    EPCPPA Error preparing to create Presidential PA
    ECPA Error creating PA (Insertion)
    CPAEGPH Create PA Error generating password
    PAEIU PA Error inserting into users
    ECP Error creating president
    ECPH Error crating presidents hash
    ECPU Error creating president user
    ECCP Errror commiting create president
    ECCPA Error coomiting create PA
    ECPRM Error creating presidential running mate
    ECRH Error creating running mate hash
    ECRU Error creating running mate user
    CAEIIVC Create Agent error inserting initial vote count to zero
    EDN Error deleting news with the looged id
    EDNZRA Error deleting news zero ros were affected (news with the looged id)
    EMNH Error marking news as handled
    EPUCV Error preparing to update candidate votes
    EUCV Error updating candidate vote count
    EUPVC Error updating presidents vote count
    EUPSVC ERROR updating polling station vote count
    ECPUVC Error commiting prsidential update vote count
    `
  fmt.Println("")
  fmt.Println(nb)
  fmt.Println("")
  fmt.Println(err)
}

func Help(){
  fmt.Println("[+] DOCUMENTATION For PRESIDENTIAL VOTER COUNTER")
  fmt.Prinln("")
  desc := `
  /*
    * All data is constructed before input no matter where it's from build it the write else you're fuckked
    * All agents represent a perticular polling station
    * Specific creators are only for creating an agent and a president
    * Transactions
    *     > Create votes or updating votes
  */`
  fmt.Println(desc)
  fmt.Println("[+] This are error stractures/ meaning as they are in your .data/logs/ files")
  fmt.Println("[+] The files are named according to the error:")
  fmt.Println("[+]    sql       Errors encountereed while communicationg with the sql server")
  fmt.Println("[+]    auth      Errors encountered during authentication or wrong login attempts")
  fmt.Println("[+]    requests  A log of all requests made to the server")
  //fmt.Println("[+]    ")
}
