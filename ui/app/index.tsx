import React, { useState } from "react";
import { createRoot } from "react-dom/client";
import {
  Box,
  Button,
  ChakraProvider,
  Heading,
  Input,
  Text,
} from "@chakra-ui/react";
import { EchoService } from "../../gen/client";

const CLIENT = new EchoService();

function App() {
  const [input, setInput] = useState<string>("");
  const [output, setOutput] = useState<string>("");
  const [err, setErr] = useState<string>("");

  return (
    <ChakraProvider>
      <Box maxWidth="80%" margin="auto" marginTop="5em">
        <Heading>Echo App</Heading>

        <Input
          onChange={(event) => setInput(event.currentTarget.value)}
          value={input}
        />
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
const root = createRoot(container!);
root.render(<App />);
