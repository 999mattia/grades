import { User } from "./types";

export function login(name: string, password: string): Promise<any> {
	return fetch("https://api.grades.mattiag.ch/api/login", {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ name, password }),
	}).then((res) => res.json());
}

export function get(id: number, token: string): Promise<User> {
	return fetch(`https://api.grades.mattiag.ch/api/user/${id}`, {
		method: "GET",
		headers: { "Content-Type": "application/json", Authorization: token },
	}).then((res) => res.json());
}
