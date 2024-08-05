function analyzeMessage() {
    const apiUrl = 'https://comdotcomdotcomdotcomdotcomdotcom.com/analyzeMessage';
    let data = {
    prompt: document.getElementById("entrybox").value,
    };

    const requestOptions = {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        // Token for backend API -- limited access
        'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMjIwOTQ2NTZ9.JCuVTr7XjdAoRpy1mO2vsy4Vnl_XLz4veVvMbif-7Wg'
    },
    body: JSON.stringify(data),
    };

    fetch(apiUrl, requestOptions)
    .then((response) => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        } return response.json();
    })
    .then((responseData) => {
        console.log(responseData);
        document.getElementById("sentiment").innerHTML = responseData.sentiment;
        document.getElementById("rephrased_statement").innerHTML = responseData.rephrased_statement;
        document.getElementById("benefits").innerHTML = JSON.stringify(responseData, null, 2);
        return responseData;
    })
    .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
    });
}