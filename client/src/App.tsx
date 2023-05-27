import React from "react";
import "./styles/App.css";
import { useNavigate } from "react-router-dom";

function App() {
	const navigate = useNavigate();
	const [loggedIn, setLoggedIn] = React.useState(false);

	React.useEffect(() => {
		const interval = setInterval(() => {
			console.log("Checking if logged in...");
			if (!localStorage.getItem("token")) {
				navigate("/login");
			}
		}, 1000);
		return () => clearInterval(interval);
	}, []);

	return <>Home</>;
}

export default App;
