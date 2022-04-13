function barista(coffees) {
  coffees.sort((a, b) => parseInt(a, 10) - parseInt(b, 10));
  let sum = coffees[0];
  let finalSum = sum || 0;
  for (let index = 1; index < coffees.length; index++) {
    sum += 2 + coffees[index];
    finalSum += sum;
  }
  return finalSum;
}

console.log(barista([])) // 0
console.log(barista([2, 10, 5, 3, 9])) // 85
console.log(barista([4, 3, 2])) // 22
console.log(barista([20, 5])) // 32
console.log(barista([20,5,4,3,1,5,7,8])) // 211
console.log(barista([5,4,3,2,1])) // 55