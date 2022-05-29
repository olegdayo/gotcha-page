import logo from './logo.svg';
import './App.css';
import CheckBoxes from "./elements/CheckBoxes";

function postRequest() {
}

function App() {
    return (
        <>
            <div className="user-info">
                <label>
                    Nickname:
                </label>
                <input type="text" name="nickname"/>
                <br/>
                <button onClick={postRequest}>Get pages!</button>
            </div>
            <CheckBoxes/>
        </>
    );
}

export default App;
