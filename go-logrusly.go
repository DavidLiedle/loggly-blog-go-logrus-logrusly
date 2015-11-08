package main

import (
    "github.com/Sirupsen/logrus"
    "github.com/sebest/logrusly"
)

/* Note: From https://github.com/Sirupsen/logrus
 * Please note the Logrus API is not yet stable (pre 1.0).
 * Logrus itself is completely stable and has been used in
 * many large deployments. The core API is unlikely to change
 * much but please version control your Logrus to make sure you
 * aren't fetching latest master on every build.
 */

var logglyToken string = "YOUR_LOGGLY_TOKEN"

func main() {
    log := logrus.New()
    hook := logrusly.NewLogglyHook(logglyToken, "https://logs-01.loggly.com/bulk/", logrus.WarnLevel, "go", "logrus")
    log.Hooks.Add(hook)

    log.WithFields(logrus.Fields{
        "name": "Slartibartfast",
        "age":  42,
    }).Error("Hello Dent, Arthur Dent!")

    // Flush is automatic for panic/fatal
    // Just make sure to Flush() before exiting or you may lose up to 5 seconds
    // worth of messages.
    hook.Flush()
}
