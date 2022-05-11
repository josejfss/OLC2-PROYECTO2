fn main(){      
     
    let xy :[[i64; 10];5    ] = [[ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11], 
    [ 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22], 
    [ 31, 32, 33, 34, 35, 36, 37, 38, 39, 30, 41],
    [ 42, 43, 53, 54, 55, 56, 57, 58, 59, 60, 61],
    [ 71, 72, 73, 74, 75, 76, 77, 78, 79, 110, 111],
    [ 81, 82, 83, 84, 85, 86, 87, 88, 89, 100, 101]];
    println!("Todo arreglo 2dim: {:?}",xy);  
    println!("Posicion [2][0] arreglo 2dim: {}",xy[1][2]);      
}

/*

let mut xd = prueba{id1:25,id2:5.0,id3:"hola"};
    println!("valor struct prueba id3: {}",xd.id3);
    println!("valor struct prueba id1: {}",xd.id1);
    xd.id1=35;
    println!("valor cambiado struct prueba id1: {}",xd.id1);
    xd.id3="adios";
    println!("valor cambiado struct prueba id3: {}",xd.id3);

    let mut y: Vec<&str> = Vec::with_capacity(1);
    y.insert(0,"Hoy");
    y.push("Si"); 
    y.push("Sale");
    y.push("Compi");
    y.push("Porfi");
    y.push("Aux");
    println!("{:?}", y);  
    y.remove(1);
    println!("{}", y.len()); 
    println!("{:?}", y);  
    println!("Entrando a funcion1");
    funcion1(5,10);

    println!("Entrando a funcion2");
    funcion2()

    let x :[i64;5] = [1,2,3,4,5];       
    println!("Todo arreglo 1dim: {:?}",x);
    println!("Posicion [3] arreglo 1dim: {}",x[3]);    
    let xy :[[i64; 4];2] = [[ 1, 3, 5, 7], [ 9, 11, 13, 15]];
    println!("Todo arreglo 2dim: {:?}",xy);  
    println!("Posicion [2][0] arreglo 2dim: {}",xy[2][0]);  
    let xyz :[[[i64; 4];2]; 2] = [
        [ [ 1, 1000, 5, 7], [ 9, 11, 13, 15] ],
        [ [ 2, 4, 6, 8], [10, 12, 14, 16] ]
        ];
    println!("Todo arreglo 3dim: {:?}",xyz);  
    println!("Posicion [1][0][0] 3dim: {}",xyz[1][0][0]); 




fn main(){    
    let mut i :i64 = 10*5+6-8/4*5;
    println!("el resultado es: {}",i);
}
Struct prueba{
    id1 :i64,
    id2 :f64,
    id3 :&str
}

fn holas(h : mut i64, o : mut i64)=>i64{
    let mut i :i64 = 0;
    println!("estoy afuera del loop: {}",h+o);
    loop{
        println!("estoy en el loop: {}",i);
        if i>5{
            break;
        }
        i=i+1;
    }
    println!("estoy afuera del loop: {}",i);
}


fn main(){      
    let mut xd = prueba{id1:25,id2:5.0,id3:"hola"};
    println!("valor struct prueba: {}",xd.id3);
    println!("valor struct prueba: {}",xd.id1);
    xd.id1=35;
    println!("valor cambiado struct prueba: {}",xd.id1);
    xd.id3="adios";
    println!("valor cambiado struct prueba: {}",xd.id3);

    let mut y: Vec<&str> = Vec::with_capacity(1);
    y.insert(0,"dem");
    y.push("hol"); 
    y.push("soy");
    y.push("muy");
    y.push("cap");
    y.push("deg");
    y.remove(1);
    println!("{}", y.len()); 
    println!("{:?}", y);  
    holas(5,10);
     
    let x :[i64;5] = [1,2,3,4,5];       
    println!("wenas: {:?}",x);  
    let xy :[[i64; 4];2] = [[ 1, 3, 5, 7], [ 9, 11, 13, 15]];
    println!("wenas bidi: {}",xy[2][0]);    
    let xyz :[[[i64; 4];2]; 2] = [
        [ [ 1, 3, 5, 7], [ 9, 11, 13, 15] ],
        [ [ 2, 4, 6, 8], [10, 12, 14, 16] ]
        ];
    println!("wenas multi: {}",xyz[1][0][0]); 
    
}

 

*/




/*


fn holas(h : mut i64, o : mut i64){
    let i :i64 = 0;
    
    loop{
        println!("estoy en el while: {}",i);
        if i>5{
            break;
        }
        i=i+1;
    }
    println!("estoy afuera del while: {}",i);
}
let xy :[[[i64; 4];2]; 2] = [
        [ [ 1, 3, 5, 7], [ 9, 11, 13, 15] ],
        [ [ 2, 4, 6, 8], [10, 12, 14, 16] ]
        ];

println!("vamooooos1234{}",h+o+i);

let ho2 :&str = "holap";
    println!("wenas2 {}",ho2);
    if true{
        let ho :i64 = -i64::pow(2,6).abs.sqrt;
        println!("prueba1 {}",ho);
        let hlo :i64 = 7%2;
        println!("estoy en el if2: {}",hlo);
    }
    ho2 = "soy lo maximo jijijiijij";
println!("wenas2 {}",ho2);
    println!("estoy fuera del if: {} kpedo","puta vida");

match 8{
        1 | 2 | 3 | 5=>{
            println!("estoy en match1");
        }
        8 =>{
            println!("estoy en match2");
        }
        _ =>{
            println!("estoy en default");
        }
    }
    loop{
        println!("wenaas");
        if true{
            println!("wenaas12");
        }
        break;
    }
    let mut x: Vec<i64> = Vec::new();
    let mut i :i64=-5+25.abs;
    println!("potencia {}",i64::pow(5,6));
    println!("abs {}", i);
*/

