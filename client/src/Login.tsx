import React from "react";
import logo from "./logo.svg";
import "./styles/App.css";
import { useNavigate } from "react-router-dom";

function Login() {
	const navigate = useNavigate();
	return (
		<>
			Login
			<button
				onClick={() => {
					localStorage.setItem("token", "test");
					navigate("/");
				}}
			>
				login
			</button>
		</>
	);
}

export default Login;
