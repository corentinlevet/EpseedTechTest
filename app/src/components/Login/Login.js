import React, { useState } from "react";
import goServer from "../../api/go-server";

import Alert from "react-bootstrap/Alert";

import './Login.css'

import userIcon from "./assets/person.png";
import emailIcon from "./assets/email.png";
import passwordIcon from "./assets/password.png";
import eyeOpenIcon from "./assets/eye-see.png";
import eyeClosedIcon from "./assets/eye-dont-see.png";

function signUpButtons({ setAction, setToken, setShowAlert, setAlertMessage, username, email, password }) {
  const signUp = async (username, email, password) => {
    if (username === "" || email === "" || password === "") {
      setShowAlert(true);
      setAlertMessage("Please fill all the fields !");
      return;
    }

    goServer.signUp(username, email, password).then((response) => {
      setToken(response.data.token);
      setShowAlert(false);
      setAlertMessage("");
      localStorage.setItem("token", response.data.token);
      localStorage.setItem("current_user_id", response.data.user_id);
    }).catch((error) => {
      console.log("Error : " + error);

      setToken("");
      setShowAlert(true);
      setAlertMessage("Error while creating the account !");
      localStorage.setItem("token", "");
      localStorage.setItem("current_user_id", "");
    });
  }

  return (
    <div className="submit-container">
      <div className="submit" onClick={() => signUp(username, email, password)}>Sign Up</div>
      <div className="submit gray" onClick={() => setAction("Login")}>Login</div>
    </div>
  );
}

function loginButtons({ setAction, setToken, setShowAlert, setAlertMessage, username, password }) {
  const logIn = (username, password) => {
    if (username === "" || password === "") {
      setShowAlert(true);
      setAlertMessage("Please fill all the fields !");
      return;
    }

    goServer.logIn(username, password).then((response) => {
      setToken(response.data.token);
      setShowAlert(false);
      setAlertMessage("");
      localStorage.setItem("token", response.data.token);
      localStorage.setItem("current_user_id", response.data.user_id);
    }).catch((error) => {
      console.log("Error : " + error);

      setToken("");
      setShowAlert(true);
      setAlertMessage("Error while logging in !");
      localStorage.setItem("token", "");
      localStorage.setItem("current_user_id", "");
    });
  }

  return (
    <div className="submit-container">
      <div className="submit gray" onClick={() => setAction("Sign Up")}>Sign Up</div>
      <div className="submit" onClick={() => logIn(username, password)}>Login</div>
    </div>
  );
}

function Login({ setToken }) {
  const [action, setAction] = useState("Sign Up");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);

  const [showAlert, setShowAlert] = useState(false);
  const [alertMessage, setAlertMessage] = useState("");

  return (
    <div className="login-screen">
      <div className="login-container">
        <div className="login-header">
          <div className="text">{action}</div>
          <div className="underline"></div>
        </div>
        <div className="inputs">
          <div className="input">
            <img src={userIcon} alt="User Icon" />
            <input type="text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} />
          </div>
          {action === "Sign Up" && (
            <div className="input">
              <img src={emailIcon} alt="Email Icon" />
              <input type="text" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
            </div>
          )}
          <div className="input">
            <img src={passwordIcon} alt="Password Icon" />
            <input type={showPassword ? "text" : "password"} placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
            <img src={showPassword ? eyeOpenIcon : eyeClosedIcon} alt="Toggle Password Visibility" onClick={() => setShowPassword(!showPassword)} width={30} height={25} />
          </div>
        </div>
        {showAlert && (
          <div style={{ display: "flex", justifyContent: "center", paddingTop: "20px" }}>
            <Alert variant="danger" onClose={() => setShowAlert(false)} dismissible className="col-md-10">
              <Alert.Heading>
                Error while {action === "Sign Up" ? "signing up" : "logging in"}:
              </Alert.Heading>
              <p>{alertMessage}</p>
            </Alert>
          </div>
        )}
        {action === "Sign Up"
          ? signUpButtons({ setAction, setToken, setShowAlert, setAlertMessage, username, email, password })
          : loginButtons({ setAction, setToken, setShowAlert, setAlertMessage, username, password })}
      </div>
    </div>
  );
}

export default Login;
