import { Box, Button, ChakraProvider, Heading, Input, Text } from "@chakra-ui/react";

import React, { useState } from "react";
import { createRoot } from "react-dom/client";

import { EchoService } from "../pb/client";

const CLIENT = new EchoService();

function App() {
  const [input, setInput] = useState<string>("");
  const [output, setOutput] = useState<string>("");
  const [err, setErr] = useState<string>("");

  return (
    <ChakraProvider>
      <Box margin="auto" marginTop="5em" maxWidth="80%">
        <Heading>Echo App</Heading>

        <Input value={input} onChange={event => setInput(event.currentTarget.value)} />
        <Button
          onClick={async () => {
            setOutput("");
            setErr("");

            try {
              const resp = await CLIENT.v1.echoServiceEcho({
                input: input,
              });

              if (resp.ok) {
                setOutput(resp.data.output || "");
              } else {
                setErr(resp.statusText);
              }
            } catch (err) {
              setErr(err.error.message);
            }
          }}
        >
          Submit
        </Button>
        <Text>output: {output}</Text>
        <Text color="red">error: {err}</Text>
      </Box>
    </ChakraProvider>
  );
}

const container = document.getElementById("root");
if (container !== null) {
  const root = createRoot(container);
  root.render(<App />);
}
