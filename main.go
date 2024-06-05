package main
import (
	"fmt"
	"os"
        "github.com/spf13/cobra"
        "strconv"
)
func printerr(err error){
     if err != nil {
        panic(err);
     }
}
func main(){

        var versionCmd = &cobra.Command{
            Use: "-v",
            Short: "display version number",
            Run: func(cmd *cobra.Command, args []string){
                 fmt.Println("gocat ver 0.1")
            },
        }

        var binaryMode bool = false;
        
        var useCmd = &cobra.Command{
            Use: "f [file to concatinate] [truncate]",
            Short: "concatinate the specfied file with standard out, truncating binary readouts as specified",
            Args: cobra.MinimumNArgs(2),
            Run: func(cmd *cobra.Command, args []string){
	         var filenm string = args[0]
                 trunc, err := strconv.Atoi(args[1])
                 printerr(err)
                 f, err := os.Open(filenm)
                 printerr(err)         
                 fi, err := f.Stat()
                 var filecon []uint8 = make([]uint8, fi.Size())
                 amnt, err := f.Read(filecon)
                 if !binaryMode { 
                     fmt.Printf("%s\n", string(filecon[:amnt])) 
                 } else{
                    for num := range filecon{
                        if 0 == num%trunc {
                           fmt.Println()
                        }
            	        fmt.Printf("%02x ", filecon[num])
            	    }
            	    fmt.Println()
            	 }
                 f.Close()
            },
        }
        


        var rootCmd = &cobra.Command{Use: "gocat"}
        rootCmd.AddCommand(versionCmd, useCmd)
        rootCmd.PersistentFlags().BoolVarP(&binaryMode, "binary", "b", false, "read file as binay and display as HEX");
        rootCmd.Execute()
}
