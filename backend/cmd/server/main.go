package main

import (
	"os"
	"log"
	"github.com/Masozee/researchopera/internal/domain/tenant"
	"github.com/Masozee/researchopera/internal/domain/user"
	"github.com/Masozee/researchopera/internal/domain/project"
	"github.com/Masozee/researchopera/internal/domain/publication"
	"github.com/Masozee/researchopera/internal/domain/dataset"
	"github.com/Masozee/researchopera/internal/domain/literature"
	"github.com/Masozee/researchopera/internal/domain/ethics"
	"github.com/Masozee/researchopera/internal/domain/grant"
	"github.com/Masozee/researchopera/internal/domain/sponsor"
	"github.com/Masozee/researchopera/internal/domain/compliance"
	"github.com/Masozee/researchopera/internal/domain/safety"
	"github.com/Masozee/researchopera/internal/domain/integrity"
	"github.com/Masozee/researchopera/internal/domain/audit"
	"github.com/Masozee/researchopera/internal/util"
)

func main() {
	log.Println("Research Opera Backend Server starting...")
	dsn := os.Getenv("MYSQL_DSN")
	db, err := util.InitDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&tenant.Tenant{},
		&user.User{},
		&project.Project{}, &project.ProjectUser{},
		&publication.Publication{}, &publication.PublicationAuthor{},
		&dataset.Dataset{},
		&literature.Literature{},
		&ethics.EthicsApproval{},
		&grant.GrantApplication{},
		&sponsor.Sponsor{},
		&compliance.Policy{}, &compliance.PolicyAcceptance{},
		&safety.IncidentReport{},
		&integrity.IntegrityDeclaration{}, &integrity.ConflictOfInterest{},
		&audit.AuditChecklist{}, &audit.AuditItem{},
	)

	// Start Fiber API server and inject DB
	api.StartServer(db)
}


