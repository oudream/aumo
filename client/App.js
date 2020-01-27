import React from "react"
import {
  ApplicationProvider,
  Layout,
  Text,
  IconRegistry,
  Button,
  Icon
} from "@ui-kitten/components"
import { mapping } from "@eva-design/eva"
import { EvaIconsPack } from "@ui-kitten/eva-icons"
import theme from "./theme"
import customM from "./mapping"

export const FacebookIcon = style => <Icon name="facebook" {...style} />

export const LoginButton = () => (
  <Button icon={FacebookIcon}>Login with Facebook</Button>
)

const HomeScreen = () => (
  <Layout style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
    <Text category="h1">Home</Text>
    <LoginButton />
  </Layout>
)

const App = () => (
  <>
    <IconRegistry icons={EvaIconsPack} />
    <ApplicationProvider
      mapping={mapping}
      theme={theme}
      customMapping={customM}
    >
      <HomeScreen />
    </ApplicationProvider>
  </>
)

export default App
