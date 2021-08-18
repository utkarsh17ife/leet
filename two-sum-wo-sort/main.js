// https://leetcode.com/problems/two-sum/ (for non sorted array)

function twoSum(nums, target) {

	// make hash table
	let ht = {}
	for (let i = 0; i < nums.length; i++) {
		ht[nums[i]] = i
	}

	let a = -1
	let b = -1

	// start from i and look for other number
	for (let i = 0; i < nums.length-1; i++) {

		let otherNum = target - nums[i]

		let j = ht[otherNum]
        if (j){
            if (i != j) {
				a = i
				b = j
				break
			}
        }
	}

	return [a, b]

}

function main() {
	let nums = [3,3]
	let res = twoSum(nums, 6)
	console.log(res[0], res[1])
}

main()