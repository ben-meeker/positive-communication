function printSc() {
    fetch("136.36.170.25:3056", {
        method: "POST",
        body: JSON.stringify({
          prompt: "What is even that",
        }),
        headers: {
          "Content-type": "application/json; charset=UTF-8",
          "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMjIwOTQ2NTZ9.JCuVTr7XjdAoRpy1mO2vsy4Vnl_XLz4veVvMbif-7Wg"
        }
      })
        .then((response) => response.json())
        .then((json) => console.log(json));

        return json
}