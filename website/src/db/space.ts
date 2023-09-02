import { noteList } from './note'

import { nanoid } from 'nanoid'

const spaces = [
  { id: nanoid(), name: 'Blog Posts' },
  { id: nanoid(), name: 'Hanna, The girl I love' },
  { id: nanoid(), name: 'More College Stuff' },
  { id: nanoid(), name: 'High School Stuff' },
  { id: nanoid(), name: 'Adventure of my live, vol.4' },
  { id: nanoid(), name: 'Adventure of my live, vol.3' },
  { id: nanoid(), name: 'Adventure of my live, vol.2' },
  { id: nanoid(), name: 'Adventure of my live, vol.1' },
  { id: nanoid(), name: 'Night time thought' },
  { id: nanoid(), name: 'Random Thoughts' },
  { id: nanoid(), name: 'Gym Bros Advice' },
  { id: nanoid(), name: 'Muscle Wiki I should write' },
  { id: nanoid(), name: 'All About today I learned' },
]

export const spaceList = spaces.map((item) => ({ ...item, notes: noteList }))
