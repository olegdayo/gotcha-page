import './CheckBoxes.css';
import axios from 'axios';
import {useEffect, useState} from "react";

function CheckBoxes() {
    const [checkboxes, setCheckboxes] = useState();
    useEffect(() => {
        axios.get("http://localhost:8080")
            .then((resp) => {
                setCheckboxes(resp.data);
            });
    }, []);
    return (
        <>
            {JSON.stringify(checkboxes)}
        </>
    );
}

export default CheckBoxes;
