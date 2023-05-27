export function login(name: string, password: string): Promise<any> {
	return fetch("https://api.grades.mattiag.ch/api/login", {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ name, password }),
	}).then((res) => res.json());
}
