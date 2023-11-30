// filename FaktorIstocasnosti.go

package main

import "fmt"

// new map type
type mmap map[int]float64


// function map.put(key, value) : insert f_* for n*
func (m mmap) put(num int, fakt float64) {
    m[num] = fakt
}

// function map.get(key) : get f_* according to n*
func (m mmap) get(num int) float64 {
    return m[num]
}

// main
func main() {

    // stedilniki
    // populate f_* according to n*
    f_st := make(mmap)
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
    f_pg := make(mmap)
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

    // real data
    ogv := 2.7      // OGV:     gas consumption for heater in [m3/h]
    nogv := 11      // nogv:    number of gas heaters
    st := 1.2       // ST:      gas consumption of cooker in [m3/h]
    nst := 11       // nst:     number of gas cookers
    
    fmt.Println()
    fmt.Printf("OGV = %.1f m3/h, število: %d --> faktor istočasnosti f(OGV) = %.3f\n", ogv, nogv, f_pg.get(nogv))
    fmt.Printf("ST = %.1f m3/h, število: %d --> faktor istočasnosti f(ST) = %.3f\n", st, nst, f_st.get(nst));
    
    poraba := (ogv * float64(nogv) * f_pg[nogv] + (st * float64(nst) * f_st[nst]))
    
    fmt.Printf("Vp = %.3f m3/h * %d * %.3f + %.3f m3/h * %d * %.3f = %.3f m3/h\n", ogv, nogv, f_pg.get(nogv), st, nst, f_st.get(nst), poraba)

}
