import React from "react";
import "./styles/App.css";
import { useNavigate } from "react-router-dom";
import { get } from "./lib/fetch";
import jwt from "jsonwebtoken";
import jwtDecode from "jwt-decode";
import { User } from "./lib/types";
import { useState, useEffect } from "react";

const defaultModel: User = {
	id: 0,
	name: "",
	modules: [],
};

function App() {
	const navigate = useNavigate();
	const [userData, setUserData] = useState<User>(defaultModel);

	async function getData() {
		let decoded: any = await jwtDecode(
			localStorage.getItem("token") as string
		);

		const res: User = await get(
			decoded.sub,
			localStorage.getItem("token") as string
		);

		setUserData(res);
	}

	useEffect(() => {
		console.log(userData);
	}, [userData]);

	useEffect(() => {
		getData();
	}, [localStorage.getItem("token")]);

	React.useEffect(() => {
		const interval = setInterval(() => {
			if (!localStorage.getItem("token")) {
				navigate("/login");
			}
		}, 500);
		return () => clearInterval(interval);
	}, []);

	return <>Home</>;
}

export default App;
