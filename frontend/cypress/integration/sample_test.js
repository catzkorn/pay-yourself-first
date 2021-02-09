describe('A user inputting values for a month', () => {
  it('saves the values, and persists them through date change', () => {

    const incomeSource = "Cat Cabin Cottages"
    const incomeAmount = 5200.00
    const savingPercent = 25
    const expensesSource = "Train Commute"
    const expensesAmount = 55

    cy.visit('http://localhost:5000')

    cy.get('#budget-date-selection')
      .get('#date-month')
      .select('March')
    cy.get('#budget-dashboard-col-align > :nth-child(1) > h5')
      .contains('03/2021')

    cy.get('#budget-date-selection')
      .get('#date-year')
      .select('2020')
      cy.get('#budget-dashboard-col-align > :nth-child(1) > h5')
        .contains('03/2020')

    cy.get('#income-source')
      .clear()
      .type(incomeSource)

    cy.get('#income-amount')
      .clear()
      .type(incomeAmount)

    cy.get('#budget-dashboard-income > :nth-child(2)')
    .contains(incomeAmount)

    cy.get('#saving-percent')
      .clear()
      .type(savingPercent)

    cy.get('#budget-dashboard-saving > :nth-child(4)')
      .contains(savingPercent + '%')

    cy.get('#expenses-table > tbody > :nth-child(1) > input')
      .clear()
      .type(expensesSource)

    cy.get('#expenses-table > tbody > :nth-child(2) > input')
    .clear()
      .type(expensesAmount)

    cy.get(':nth-child(3) > select')
      .select('Monthly')

    cy.get('#budget-dashboard-income > :nth-child(4)')
      .contains(expensesAmount)

    cy.get('button')
      .click()
  })
})