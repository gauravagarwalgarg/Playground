# Product Engineering

## Methodologies

### Agile
- **Values**: Individuals over processes, working software over documentation, collaboration over contracts, responding to change over following a plan
- **Principles**: Deliver frequently, welcome change, sustainable pace, technical excellence
- **Not**: No planning, no documentation, no deadlines

### Scrum
- **Roles**: Product Owner, Scrum Master, Development Team
- **Ceremonies**: Sprint Planning, Daily Standup, Sprint Review, Retrospective
- **Artifacts**: Product Backlog, Sprint Backlog, Increment
- **Sprint**: Fixed time-box (1-4 weeks), shippable increment at end

### Kanban
- **Principles**: Visualize work, limit WIP, manage flow, make policies explicit
- **Metrics**: Lead time, cycle time, throughput
- **When**: Continuous flow, support/ops teams, unpredictable work

### CI/CD (Continuous Integration / Continuous Delivery)
- **CI**: Merge frequently, automated build + test on every push
- **CD (Delivery)**: Code is always in a deployable state
- **CD (Deployment)**: Every change that passes pipeline goes to production
- **Pipeline**: Build → Unit Tests → Integration Tests → Deploy to Staging → Deploy to Prod
- **Tools**: GitHub Actions, Jenkins, GitLab CI, ArgoCD

### DevOps Culture
- **Principles**: Shared ownership, automation, measurement, feedback loops
- **Practices**: Infrastructure as Code, monitoring, incident response, blameless postmortems
- **SRE overlap**: SLOs, error budgets, toil reduction
- **Key metrics (DORA)**: Deployment frequency, lead time, change failure rate, MTTR

---

## Interview Talking Points

- Agile is a mindset, not a process new_textScrum/Kanban are frameworks
- CI/CD reduces risk by making deployments small and frequent
- DevOps breaks down silos between dev and ops new_text"you build it, you run it"
- Measure what matters: DORA metrics correlate with high-performing teams
