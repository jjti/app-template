import React from "react";
import ReactDOM from "react-dom";
import { ChakraProvider } from "@chakra-ui/react";

function App() {
  return (
    <ChakraProvider>
      <div>Hello, esbuild! veriosn 5</div>
    </ChakraProvider>
  );
}

ReactDOM.render(<App />, document.getElementById("root"));
