import './CheckBoxes.css';
import axios from 'axios';

async function CheckBoxes() {
    let response;
    response = await axios.get("http://localhost:8080")
        .catch(function (error) {
            console.error(error);
        });
    console.log(response);
    console.log(response.data);
    let checkboxes = response.data;
    console.log(checkboxes);
    return (
        <>
        </>
    );
}

export default CheckBoxes;
