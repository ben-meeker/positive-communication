function printSc() {
    const apiUrl = 'https://comdotcomdotcomdotcomdotcomdotcom.com/analyzeMessage';
    let data = {
    "prompt": 'What is even that',
    };

    const requestOptions = {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMjIwOTQ2NTZ9.JCuVTr7XjdAoRpy1mO2vsy4Vnl_XLz4veVvMbif-7Wg'
    },
    };
    //body: JSON.stringify(data),
    //};

    fetch(apiUrl, requestOptions)
    .then(response => {
        if (!response.ok) {
        throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        // outputElement.textContent = JSON.stringify(data, null, 2);
        return data
    })
    .catch(error => {
        console.error

    ('Error:', error);
    });

    return data
}