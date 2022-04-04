function Xbonacci(signature, n) { 
  if (n === 0) {
    return []
  }
  
  let xBonacchiLength = signature.length - 1;
  let outputSequence = [...signature]
  
  for (let index = 0; index < n - signature.length; index++) {
    let nextNumber = 0;
    
    for (let num = xBonacchiLength + index; num >= index; num--) {
      nextNumber += outputSequence[num];
    }
    outputSequence.push(nextNumber);

  }
  
  return outputSequence.slice(0,n);
}

console.log(Xbonacci([ 0, 0, 0, 0, 1 ] ,10)) //[0, 0, 0, 0, 1, 1, 2, 4, 8, 16]