// Debelina_stene.go

package main

import (
    "fmt"
    "os"
    "os/exec" // for cls()
    "strconv"
)

    var Tmin float64    // debelina stene [mm]
    var D float64       // zunanji premer cevi [mm]
    var DP float64      // delovni tlak [bar]
    var f0 float64      // faktor rabe [-]
    var Sgm_p float64   // obodna napetost [N/mm2]
    var Rt float64      // nanjnižja specifična meja plastičnosti [N/mm2]

// clear the screen on cmd
func Cls_win() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// clear the screen on *nix
func Cls_nix() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func Usage() {
    fmt.Println(
    "\nUsage: Debelina_stene[.exe] <command argumets>" + "\n")
    
    fmt.Print(
    `   command arguments: Rt f0 DP D
            Rt - nanjnižja specifična meja plastičnosti [N/mm2]
            f0 - faktor rabe [-]
            DP - delovni tlak [bar]
            D  - zunanji premer cevi [mm]`)
    fmt.Println("\n\n")

}

func Izracun_na_faktor(Faktor float64) float64 {
    Sgm_p = Faktor * Rt
    Tmin = (DP * D) / (20 * Sgm_p)
    return Tmin
}

func main() {
    
    Cls_win()
    Cls_nix()

    if len(os.Args) != 5 {
        Usage()
        
        Rt = 245.0
        f0 = 0.4
        DP = 50
        D = 114.3
    } else {
        Rt, _ = strconv.ParseFloat(os.Args[1], 64)
        f0, _ = strconv.ParseFloat(os.Args[2], 64)
        DP, _ = strconv.ParseFloat(os.Args[3], 64)
        D, _  = strconv.ParseFloat(os.Args[4], 64)
    }
    
    fmt.Println()
    fmt.Println("--------------------------------------------------")
    fmt.Printf("%-22s %-10s %-10s %-8s\n", "Naziv", "Oznaka", "Vrednost", "Enota")
    fmt.Println("--------------------------------------------------")
    fmt.Printf("%-22s %-10s %-10.2f %-8s\n", "Delovni tlak", "DP", DP, "bar")
    fmt.Printf("%-22s %-10s %-10.2f %-8s\n", "Premer cevi", "D", D, "mm")
    fmt.Printf("%-22s %-10s %-10.2f %-8s\n", "Meja plastičnosti", "Rt",  Rt, "N/mm2")
    fmt.Printf("%-22s %-10s %-10.2f %-8s\n", "Faktor", "f0", f0, "-")
    fmt.Println("--------------------------------------------------")
    fmt.Printf("%-22s %-10s %-10.2f %-8s\n", "Min. debeliba stene", "Tmin", Izracun_na_faktor(f0), "mm")
    fmt.Println("==================================================")
    
    F0 := [3]float64{0.4, 0.5, 0.6} 
    
    for _, Fk := range(F0) {
        fmt.Printf("Min. debelina stene Tmin[f0 = %.1f] = %.2f mm\n", Fk, Izracun_na_faktor(Fk))
    }

}


