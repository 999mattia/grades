import "./styles/App.css";
import { useNavigate } from "react-router-dom";
import { login } from "./lib/fetch";
import { useState } from "react";

const defaultModel = {
	name: "",
	password: "",
};

function Login() {
	const navigate = useNavigate();
	const [user, setUser] = useState(defaultModel);

	const handleChange = (e: any) => {
		const target = e.target;
		const name = target.name;
		const value = target.value;
		setUser({
			...user,
			[name]: value,
		});
	};

	const handleLogin = async () => {
		const response = await login(user.name, user.password);

		localStorage.setItem("token", response.token);

		setUser(defaultModel);

		navigate("/");
	};

	return (
		<>
			<button onClick={() => handleLogin()}>Login</button>
		</>
	);
}

export default Login;
