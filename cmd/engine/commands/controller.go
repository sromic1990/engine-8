package commands

import (
	"github.com/battlesnakeio/engine/controller/filestore"
	log "github.com/sirupsen/logrus"

	"github.com/battlesnakeio/engine/controller"
	"github.com/spf13/cobra"
)

var (
	controllerListen      = ":3004"
	controllerBackend     = "inmem"
	controllerBackendArgs = ""
)

func init() {
	controllerCmd.Flags().StringVarP(&controllerListen, "listen", "l", controllerListen, "address for the controller to bind to")
	controllerCmd.Flags().StringVarP(&controllerBackend, "backend", "b", controllerBackend, "controller backend, as one of: [inmem, file]")
	controllerCmd.Flags().StringVarP(&controllerBackendArgs, "backend-args", "a", controllerBackendArgs, "options to pass to the backend being used")
	allCmd.Flags().AddFlagSet(controllerCmd.Flags())
}

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "runs the engine controller",
	Run: func(c *cobra.Command, args []string) {
		var store controller.Store
		switch controllerBackend {
		case "inmem":
			store = controller.InMemStore()
		case "file":
			store = filestore.NewFileStore(controllerBackendArgs)
		default:
			log.WithField("backend", controllerBackend).Fatal("invalid backend")
		}

		ctrl := controller.New(store)
		log.WithField("listen", controllerListen).
			Info("Battlesnake controller serving")
		if err := ctrl.Serve(controllerListen); err != nil {
			log.WithError(err).
				WithField("listen", controllerListen).
				Fatal("Controller failed to serve")
		}
	},
}