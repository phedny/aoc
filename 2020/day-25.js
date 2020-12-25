function findSecretLoopSize(publicKey) {
	let value = 1;
	for (let i = 0; ; i++) {
		if (value === publicKey) {
			return i;
		}
		value *= 7;
		value %= 20201227;
	}
}

function transform(loopSize, num) {
	let value = 1;
	for (let i = 0; i < loopSize; i++) {
		value *= num;
		value %= 20201227;
	}
	return value;
}

// const [public1, public2] = [5764801, 17807724];
const [public1, public2] = [8184785, 5293040];

const private1 = findSecretLoopSize(public1);
const result = transform(private1, public2);

console.log(result);
