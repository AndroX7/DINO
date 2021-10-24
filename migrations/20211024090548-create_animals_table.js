'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    /**
     * Add altering commands here.
     *
     * Example:
     * await queryInterface.createTable('users', { id: Sequelize.INTEGER });
     */
    await queryInterface.createTable('animals',{
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER
      },
      animal_type: {
        allowNull: false,
        type: Sequelize.STRING
      },
      nickname: {
        allowNull: false,
        unique:true,
        type: Sequelize.STRING
      },
      zone: {
        allowNull: false,
        type: Sequelize.INTEGER
      },
      age: {
        allowNull: false,
        type: Sequelize.INTEGER
      },
    })
  },

  down: async (queryInterface, Sequelize) => {
    /**
     * Add reverting commands here.
     *
     * Example:
     * await queryInterface.dropTable('users');
     */
     await queryInterface.dropTable('animals');
  }
};
