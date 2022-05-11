
fn main() {
    
    println!("-----------------------------------");
    println!("----ARCHIVO BASICO SALIDA EN 3D----");
    println!("-----------------------------------");
        
    
    // let bo1: bool = false;
    // let bol21: bool = !bo1;
    let cad1: String = "imprimir".to_string();
    let cad21: String = "cadena valida".to_string();
    let letra1: char = 'c';
    
    
    let val11 = 10 - (5 + 10 * (9 + 4 * (1 + 2 * 3)) - 8 * 3 * 6) + 5 * (2 * 2);
    let val21 = (2 * 9 * 2 * 2) - 9 - (8 - 16 + (3 * 3 - 6 * 5 - 7 - (9 + 7 * 7 * 7) + 10) - 5) + 8 - (6 - 5 * (2 * 3));
    let val31 = val11 + ((2 + val21 * 3) + 1 - ((2 * 2 * 2) - 2) * 2) - 2;
	
	
    
    let a:i64 = 100;
    let b:i64 = 100;
    let c:i64 = 7;
    // // let f:bool = true;
    let j:f64 = 10.0;
    let k:f64 = 10.0;
    
	
    let val1:i64 = 5;
    let resp:i64 = 5;
    let mut valorVerdadero : i64 = 100;
	
	
    let x1: i64 = 15;
	
	


    let abs1:i64 = 7-11.abs;
    let abs2:f64 = 10.0.abs;
    let raiz1:i64 = 9;
    let raiz2:f64 = 100.0;
    
    println!("");
    println!("*************PRUEBA DE NATIVAS");
    println!(" valor de b: {}",b);
    
    println!(" valor absoluto1: {}",abs1);
    println!(" valor absoluto2: {}",abs2);

    
    println!("El valor de val11 es:              {}",val11);
    println!("El valor de val21 es:              {}",val21);
    println!("El valor de val31 es:              {}",val31);
    println!("El resultado de la operación es:  {}",val31);
    println!("El valor de cad1 es:               {}",cad1);
    println!("El valor de cad21 es:               {}",cad21);
    println!("El valor de letra1 es:             {}",letra1);
    println!("");
    

    println!("");
    println!("");
    if a > b || b < c {
        println!(">>>>>> Esto no debería de imprimirse");
    }else{
        println!(">>>>>> Esto debería de imprimirse");
    }
    
    if a == b && j == k {
        println!(">>>>>> Esto debería de imprimirse");
    }else{
        println!(">>>>>> Esto no debería de imprimirse");
    }

    let com = 50 + 50 + val1 - val1;
    if valorVerdadero == com {
        println!(">>>>>> En este lugar deberia de entrar :)");
        valorVerdadero = 50;
    }
    else if valorVerdadero > 50 && resp != 100 {
        println!(">>>>>> Aca no deberia de entrar :ccc");
        valorVerdadero = 70;
    }
    else{
        println!(">>>>>> Aca no deberia de entrar :cccc");
    }

    if x1 % 2 == 0 {
        println!(">>>>> numeroPar ingreso a if verdadero, {} es par",x1);
    }
    else {
        println!(">>>>> numeroPar ingreso a if falso, {} no es par",x1);
    }
    
    
}

