// PreracunPorabeZP.go

package main

import (
    "fmt"
    "os"
    "os/exec" // for cls()
    "bufio"
    "strconv"
    "strings"
)

var ogv float64     // OGV:     gas consumption for heater in [m3/h]
var nogv int        // nogv:    number of gas heaters
var st float64      // ST:      gas consumption of cooker in [m3/h]
var nst int         // nst:     number of gas cookers
var input string
var report string


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


// new map type
type map1 map[int]float64
type map2 map[float64]int


// function map.put(key, value) : insert f_* for n*
func (m1 map1) put(num int, fakt float64) {
    m1[num] = fakt
}

// function map.get(key) : get f_* according to n*
func (m1 map1) get(num int) float64 {
    return m1[num]
}


func (m2 map2) put2(ogv float64, nogv int) {
    m2[ogv] = nogv
}

func (m2 map2) get2(ogv float64) int {
    return m2[ogv]
}

// main
func main() {

    // stedilniki
    // populate f_* according to n*
    f_st := make(map1)
    f_st.put(1, 0.621); f_st.put(2, 0.448); f_st.put(3, 0.371); f_st.put(4, 0.325); f_st.put(5, 0.294);
    f_st.put(6, 0.271); f_st.put(7, 0.253); f_st.put(8, 0.239); f_st.put(9, 0.227); f_st.put(10, 0.217);
    f_st.put(11, 0.208); f_st.put(12, 0.201); f_st.put(13, 0.194); f_st.put(14, 0.188); f_st.put(15, 0.183);
    f_st.put(16, 0.178); f_st.put(17, 0.173); f_st.put(18, 0.169); f_st.put(19, 0.166); f_st.put(20, 0.162);
    f_st.put(21, 0.159); f_st.put(22, 0.156); f_st.put(23, 0.153); f_st.put(24, 0.151); f_st.put(25, 0.148);
    f_st.put(26, 0.146); f_st.put(27, 0.144); f_st.put(28, 0.142); f_st.put(29, 0.140); f_st.put(30, 0.138);
    f_st.put(31, 0.136); f_st.put(32, 0.134); f_st.put(33, 0.133); f_st.put(34, 0.131); f_st.put(35, 0.130);
    f_st.put(36, 0.128); f_st.put(37, 0.127); f_st.put(38, 0.126); f_st.put(39, 0.125); f_st.put(40, 0.123);
    f_st.put(41, 0.122); f_st.put(42, 0.121); f_st.put(43, 0.120); f_st.put(44, 0.119); f_st.put(45, 0.118);
    f_st.put(46, 0.117); f_st.put(47, 0.116); f_st.put(48, 0.115); f_st.put(49, 0.114); f_st.put(50, 0.114);

    
    // grelniki
    // populate f_* according to n*
    f_pg := make(map1)
    f_pg.put(1, 1.000); f_pg.put(2, 0.883); f_pg.put(3, 0.822); f_pg.put(4, 0.782); f_pg.put(5, 0.752);
    f_pg.put(6, 0.729); f_pg.put(7, 0.710); f_pg.put(8, 0.694); f_pg.put(9, 0.680); f_pg.put(10, 0.668);
    f_pg.put(11, 0.657); f_pg.put(12, 0.648); f_pg.put(13, 0.639); f_pg.put(14, 0.631); f_pg.put(15, 0.624);
    f_pg.put(16, 0.617); f_pg.put(17, 0.611); f_pg.put(18, 0.605); f_pg.put(19, 0.599); f_pg.put(20, 0.594);
    f_pg.put(21, 0.590); f_pg.put(22, 0.585); f_pg.put(23, 0.581); f_pg.put(24, 0.577); f_pg.put(25, 0.573);
    f_pg.put(26, 0.569); f_pg.put(27, 0.566); f_pg.put(28, 0.562); f_pg.put(29, 0.559); f_pg.put(30, 0.556);
    f_pg.put(31, 0.553); f_pg.put(32, 0.550); f_pg.put(33, 0.547); f_pg.put(34, 0.545); f_pg.put(35, 0.542);
    f_pg.put(36, 0.540); f_pg.put(37, 0.537); f_pg.put(38, 0.535); f_pg.put(39, 0.533); f_pg.put(40, 0.530);
    f_pg.put(41, 0.528); f_pg.put(42, 0.526); f_pg.put(43, 0.524); f_pg.put(44, 0.522); f_pg.put(45, 0.520);
    f_pg.put(46, 0.518); f_pg.put(47, 0.517); f_pg.put(48, 0.515); f_pg.put(49, 0.513); f_pg.put(50, 0.512);
    
    ogvs := make(map2)
    sts :=  make(map2)

    // input
    Cls_win()
    Cls_nix()

    fmt.Print("\n" +
    ` Za vsak tip plinskega grelnik glede na porabo
 vnesi: [poraba] [število grelnikov]
 > `)
    for {
        reader := bufio.NewReader(os.Stdin)
        imput, _ := reader.ReadString('\n')
        s := strings.Fields(imput)
        if len(s) == 0 {
            break
        } else if s[0] == "q" {
            break
        } else {
            ogv, _ := strconv.ParseFloat(s[0], 64)
            nog, _ := strconv.Atoi(s[1])
            ogvs.put2(ogv, nog)
        }
    }
    // test
    // fmt.Println(ogvs)
    
    fmt.Print("\n" +
    ` Za vsak tip štedilnika glede na porabo
 vnesi: [poraba] [število štedilnikov]
 > `)
    for {
        reader := bufio.NewReader(os.Stdin)
        imput, _ := reader.ReadString('\n')
        s := strings.Fields(imput)
        if len(s) == 0 {
            break
        } else if s[0] == "" {
            break
        } else {
            stv, _ := strconv.ParseFloat(s[0], 64)
            nst, _ := strconv.Atoi(s[1])
            sts.put2(stv, nst)
        }
    }
    // test
    // fmt.Println(sts)
    
    ogvs_s := []string{}
    ogvs_v := 0.0
    for key, value := range ogvs {
        ogv_f := f_pg.get(value)
        ogvs_s = append(ogvs_s, fmt.Sprintf("%.2f * %d * %.3f", key, value, ogv_f))
        ogvs_v += key * float64(value) * ogv_f
    }
    
    ogvss := strings.Join(ogvs_s, " + ")
    
    report += fmt.Sprintf("\n" + "Skupna poraba za plinske grelnike:\n" +
    "%s = %.2f m3/h\n", ogvss, ogvs_v)
    
    
    stvs_s := []string{}
    stvs_v := 0.0
    for key, value := range sts {
        stv_f := f_st.get(value)
        stvs_s = append(stvs_s, fmt.Sprintf("%.2f * %d * %.3f", key, value, stv_f))
        stvs_v += key * float64(value) * stv_f
    }
    
    stvss := strings.Join(stvs_s, " + ")

    report += fmt.Sprintf("\n" + "Skupna poraba za plinske štedilnike:\n" +
    "%s = %.2f m3/h\n", stvss, stvs_v)
    
    sporaba := ogvs_v + stvs_v

    report += fmt.Sprintf("\n" + "Skupna vršna poraba vseh trošil: %.2f m3/h\n", sporaba)
    
    // final report
    fmt.Print(report)
    
    

}


