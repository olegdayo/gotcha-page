import './CheckBoxes.css';
import axios from 'axios';
import {useEffect, useState} from "react";

function CheckBoxes() {
    const [checkboxes, setCheckboxes] = useState([]);
    useEffect(() => {
        axios.get("http://localhost:8080")
            .then((resp) => {
                setCheckboxes(resp.data);
            })
            .catch((err) => {
                console.log(err)
            });
    }, []);

    return (
        <>
            <div className="checks">
                <ul>
                    {checkboxes.map(
                        (x, k) => <li key={k}>
                            <label form={x.id}>{x.name}</label>
                            <input type="checkbox" name={x.id} id={x.id}/>
                        </li>
                    )}
                </ul>
            </div>
        </>
    );
}

export default CheckBoxes;
