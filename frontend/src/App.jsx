import './App.css';
import React, { useEffect, useState } from "react";
import axios from "axios";

function App() {
    const [checkboxes, setCheckboxes] = useState([]);
    const [nickname, setNickname] = useState("");
    const [links, setLinks] = useState([]);
    useEffect(() => {
        axios.get("http://localhost:8080")
            .then((resp) => {
                setCheckboxes(resp.data.networks);
            })
            .catch((err) => {
                console.log(err)
            });
    }, []);

    function changeAll(event) {
        let selectAllCheckbox = checkboxes.find(x => x.id === event.target.name)
        setCheckboxes(checkboxes.map(
            (x) => {
                return {
                    ...x, value: !selectAllCheckbox.value
                }
            }
        ))
    }

    function change(event) {
        let checkbox = checkboxes.find(x => x.id === event.target.name)
        console.log(checkbox)
        console.log(event.target.name)
        checkbox.value = !checkbox.value
        setCheckboxes(checkboxes)
    }

    function postRequest() {
        axios.post("http://localhost:8080", {
                nickname: nickname,
                parsers: checkboxes.filter(x => x.value).map(
                    x => x.id
                )
            })
            .then((resp) => {
                setLinks(resp.data);
                console.log(resp.data)
                console.log(links)
            })
            .catch(
                console.log
            )
    }

    return (
        <>
            <div className="user-info">
                <label>
                    Nickname:
                </label>
                <input value={nickname} onChange={(event) => { setNickname(event.target.value) }} type="text" name="nickname"/>
                <br />
                <button onClick={postRequest}>Get pages!</button>
            </div>

            <div className="checks">
                <ul>
                    <li>
                        <input onChange={changeAll} type="checkbox" name="all" id="all" checked={"false"} />
                        <label form={"all"}>{"Select all"}</label>
                    </li>
                    {checkboxes.map(
                        (x, k) => <li key={k}>
                            <input onChange={change} type="checkbox" name={x.id} id={x.id} checked={x.value}/>
                            <label form={x.id}>{x.name}</label>
                        </li>
                    )}
                </ul>
            </div>

            <div className="links">
                <ul>
                    {
                    links.map(
                        (x, k) => {
                            console.log(x)
                            if (x.available) {
                                return  <li key={k}>
                                            <a name={x.url} href={x.link}>{x.name}</a>
                                        </li>
                            }
                            return  <li key={k}>
                                        <a>{x.name}</a>
                                    </li>
                        }
                    )}
                </ul>
            </div>
        </>
    );
}

export default App;
