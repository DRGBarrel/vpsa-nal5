//package redovanje omogoča dodajanje in izpisovanje ocen in uspeha učencev
package redovanje
import "fmt"

var stOcen = 5
var minOcena = 1
var maxOcena = 10
func SetSt(nov int){
    stOcen=nov
}
func SetMin(nov int){
    minOcena=nov
}
func SetMax(nov int){
    maxOcena=nov
}
type Student struct {
    ime     string
    priimek string
    ocene   []int
}
func AddStudent(studs map[string]Student, vpisna string, ime string, priimek string, ocene []int){
    studs[vpisna]=Student{ime,priimek,ocene}
}
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64{
    _,ok:=studenti[vpisnaStevilka]
    if(!ok) {
        return -1.0
    }
    sum:=0.0
    for _,v :=range studenti[vpisnaStevilka].ocene{
        sum+=float64(v)
    }
    if(sum/float64(len(studenti[vpisnaStevilka].ocene))<6.0){
        return 0.0
    }
    return sum/float64(len(studenti[vpisnaStevilka].ocene))
}
//DodajOceno doda oceno študentu, ki mu pripada vpisnaStevilka
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int){
    _,ok:=studenti[vpisnaStevilka]
    if(!ok) {
        fmt.Println("študenta s to vpisno številko ni")
    } else {
        ok=ocena>=minOcena&&ocena<=maxOcena
        if(!ok){
            fmt.Println("neveljavna ocena")
        } else{
            var tmp=make([]int,len(studenti[vpisnaStevilka].ocene)+1)
            for i,v := range studenti[vpisnaStevilka].ocene{
                tmp[i]=v
            }
            tmp[len(tmp)-1]=ocena
            studenti[vpisnaStevilka]=Student{studenti[vpisnaStevilka].ime,studenti[vpisnaStevilka].priimek,tmp}
        }
    }
}
//IzpisVsehOcen izpiše vse ocene vseh študentov
func IzpisVsehOcen(studenti map[string]Student){
    fmt.Println("REDOVALNICA:")
    for stev,stud:=range studenti{
        fmt.Printf("%s - %s %s: [",stev,stud.ime,stud.priimek)
        for v := range len(stud.ocene)-1{
            fmt.Printf("%d ",stud.ocene[v])
        }
        fmt.Printf("%d]\n",stud.ocene[len(stud.ocene)-1])
    }
}
//IzpisiKoncniUspeh izpiše povprečno oceno vsakega študenta in njihovo uspešnost:
//če je povprečje >=9, je študent odličen, če je povprečje >=6, je študent povprečen, drugače je neuspešen
func IzpisiKoncniUspeh(studenti map[string]Student){
    for stev,stud:=range studenti{
        povp:=povprecje(studenti,stev)
        fmt.Printf("%s %s: povprečna ocena %.1f -> ",stud.ime,stud.priimek,povp)
        if(len(stud.ocene)<stOcen){
            fmt.Println("Študent nima dovolj ocen za uspeh")
        }else{
            switch {
            case povp>=9:
                fmt.Println("Odličen študent!")
            case povp>=6:
                fmt.Println("Povprečen študent")
            default:
                fmt.Println("Neuspešen študent")
            }
        }
    }
}