function main(input){
	let nums = input.trim().split(" ");
	let conditions = ["12","23","45","56","78","89"];
	if(conditions.indexOf(nums.join(""))>=0){
		console.log("Yes");
	}else{
		console.log("No");
	}//if
} // input
main(require('fs').readFileSync('/dev/stdin', 'utf8'));