import './CheckBoxes.css';
import axios from 'axios';

function CheckBoxes() {
    let response = axios.get("http://localhost:8080");
    alert(response)
    return (
        <>
        </>
    );
}

export default CheckBoxes;
