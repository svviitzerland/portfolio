package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"

	"ssh-portfolio/internal/analytics"
	"ssh-portfolio/internal/data"
	"ssh-portfolio/internal/ui"
)

const (
	host    = "0.0.0.0"
	port    = "2222"
	version = "1.0.0"
)

func main() {
	// Load CV data
	cv, err := data.LoadCV()
	if err != nil {
		log.Fatal("Failed to load CV data", "error", err)
	}

	analytics.Init()

	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════════════╗")
	fmt.Println("  ║       SSH Portfolio - Farhan Aulianda            ║")
	fmt.Println("  ║       Powered by Wish + Bubble Tea              ║")
	fmt.Println("  ╚══════════════════════════════════════════════════╝")
	fmt.Println()

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithVersion("SSH-Portfolio-"+version),
		wish.WithMiddleware(
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				username := s.User()
				if username == "" {
					username = "anonymous"
				}
				log.Info("New visitor connected",
					"user", username,
					"remote", s.RemoteAddr().String(),
				)
				analytics.TrackVisitor(username, s.RemoteAddr().String())
				m := ui.NewModel(cv, username)
				return m, []tea.ProgramOption{}
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal("Could not create server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting SSH portfolio server",
		"host", host,
		"port", port,
		"connect", fmt.Sprintf("ssh -p %s %s", port, host),
	)
	fmt.Printf("  → Connect with: ssh -p %s localhost\n\n", port)

	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Shutting down SSH portfolio server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server gracefully", "error", err)
	}

	log.Info("Server stopped. Goodbye!")
}
