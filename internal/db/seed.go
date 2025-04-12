package db

import (
	"os"
	"fmt"
	"context"
	"math/rand"
	"path/filepath"
	"database/sql"
	"github.com/paulsonlegacy/go-social/internal/models"
	envloader "github.com/paulsonlegacy/go-env-loader"
)


// PACKAGE LEVEL VARIABLES

var(
	randomNames = []string{
		"paulson", "zoe", "david", "lucas", "emma", "oliver", "mason", "sophia", "jackson", "ava", 
		"liam", "isabella", "noah", "mia", "jacob", "charlotte", "william", "amelia", "elijah", "harper", 
		"james", "evelyn", "benjamin", "abigail", "lucy", "jack", "ella", "michael", "scarlett", "ethan", 
		"grace", "daniel", "lily", "henry", "zoey", "gabriel", "audrey", "sebastian", "chloe", "aiden", 
		"nora", "joseph", "aria", "levi", "hannah", "samuel", "addison", "logan", "brooklyn", "mateo", 
		"ellie", "dylan", "charlie", "harrison", "matthew", "sophia", "daisy", "miles", "elizabeth", "michael", 
		"jackson", "lucas", "adeline", "harry", "avery", "leo", "audrey", "eva", "ashton", "grayson", 
		"madison", "ian", "gianna", "william", "layla", "theodore", "emily", "thomas", "luna", "julia", 
		"jack", "victoria", "julian", "delilah", "david", "riley", "nathan", "luca", "maya", "caleb", 
		"lauren", "elise", "cameron", "jasmine", "owen", "catherine", "evan", "aurora", "leo", "grace", 
		"andrew", "piper", "oscar", "zoey", "ryan", "zoe", "eliza", "eliana", "andrew", "lena", 
		"hayden", "kylie", "lucas", "olivia", "henry", "zoe", "sofia", "jack", "joseph", "mila", 
		"audrey", "jack", "dylan", "lindsey", "hudson", "addison", "lucas", "anna", "alexander", "harper", 
		"isabelle", "oscar", "peter", "ariel", "sienna", "logan", "alexis", "travis", "zachary", "luke", 
		"maya", "kaitlyn", "gabriel", "anna", "sofia", "eliza", "audrey", "mason", "zoey", "finn", 
		"hayley", "jameson", "camilla", "ryan", "auguster", "luke", "harley", "james", "emilia", "claire", 
		"matthew", "grace", "natalie", "joseph", "nora", "emily", "louis", "riley", "tyler", "jack", 
		"zoe", "scarlett", "emma", "matthew", "michael", "seraphina", "oliver", "evan", "nicholas", 
		"george", "laura", "vivianna", "sarah", "adam", "lucy", "charlotte", "emma", "kaitlyn", "hammed", 
		"lucas", "grace", "wade", "isabelle", "scarlett", "sean", "jason", "maria", "zach", "lisa", "xena",
	}
	randomTitles = []string{
		"This is a title", "How to Build a Strong Web Application", "The Best Practices for Writing Clean Code", "Getting Started with Go Programming", "A Beginner's Guide to Machine Learning", "Understanding SQL Databases", 
		"Tips for Improving Your JavaScript Skills", "Top 10 ReactJS Libraries You Should Know", "Building REST APIs with Django", "How to Optimize Your Website for SEO", "Introduction to GraphQL",
		"Why Python is Great for Data Science", "Exploring the World of Cloud Computing", "Building a Portfolio Website from Scratch", "The Future of Artificial Intelligence", "Understanding the Basics of Blockchain",
		"5 Common Mistakes to Avoid in Software Development", "How to Use Docker in Your Projects", "Mastering the Art of Git and GitHub", "Creating a Chatbot with Python", "An Overview of Web Accessibility", 
		"Building Scalable Microservices with Kubernetes", "Getting Started with DevOps", "Why You Should Learn Rust", "How to Build an E-commerce Website", "The Power of Open Source Software", 
		"Exploring the Benefits of Serverless Architecture", "React vs Angular: Which Should You Choose?", "Best Practices for Writing Unit Tests", "How to Integrate Payments into Your Website", 
		"Setting Up Continuous Integration for Your Project", "Introduction to Agile Development", "The Importance of Data Privacy in Web Applications", "How to Get Started with Cybersecurity", 
		"Building an Online Blog with WordPress", "A Guide to Frontend Frameworks", "10 Tips for Writing Better API Documentation", "The Importance of Code Reviews", "How to Migrate from SQL to NoSQL Databases", 
		"Understanding the Basics of Cryptography", "Best Tools for Web Development", "How to Create a RESTful API", "The Complete Guide to Web Scraping", "What is the Internet of Things (IoT)?", 
		"How to Use Version Control in Software Projects", "The Difference Between SQL and NoSQL Databases", "A Beginner's Guide to Firebase", "Introduction to Node.js", "10 Coding Challenges to Improve Your Skills", 
		"How to Deploy Your Application to Heroku", "Setting Up a MongoDB Database", "Why Should You Learn Kubernetes?", "Introduction to Machine Learning Algorithms", "The Role of Artificial Intelligence in Business", 
		"Creating a Secure Login System with JWT", "A Complete Guide to Docker Compose", "How to Optimize Your Code for Performance", "The Importance of Test-Driven Development", "Building a Content Management System", 
		"Best Practices for REST API Design", "Introduction to Full-Stack Development", "How to Handle Errors in JavaScript", "Mastering Asynchronous Programming in JavaScript", "Why Is Agile Development Important?", 
		"How to Use WebSockets in Web Applications", "Understanding Design Patterns in Software Development", "How to Improve Your Web App's UX/UI", "Building Secure APIs", "Exploring the World of Data Engineering", 
		"How to Create Your Own Web Framework", "Best Cloud Services for Developers", "Building Scalable Web Applications", "Exploring the Latest Trends in Web Development", "How to Write Efficient SQL Queries", 
		"Building Cross-Platform Mobile Apps with React Native", "What Is Microservices Architecture?", "How to Get Started with Data Analysis", "The Best IDEs for Web Development", "Understanding Webhooks", 
		"An Introduction to Data Visualization", "The Basics of Network Security", "Why You Should Learn Rust for System Programming", "How to Make Your Web Application Fast", "The Benefits of Pair Programming", 
		"Best Practices for Using Git in Your Projects", "Creating a Simple To-Do App with React", "How to Use Redux in React Applications", "The Importance of Caching in Web Applications", "Best Resources for Learning JavaScript", 
		"Building Your First Android App", "The Best Tools for Debugging Your Code", "Understanding Memory Management in C", "How to Use GraphQL with React", "Mastering the Basics of Python Programming", 
		"How to Work with Dates and Times in JavaScript", "Understanding HTTP Requests and Responses", "Best Practices for Writing Clean and Maintainable Code", "How to Implement OAuth2 in Your Application", 
		"Getting Started with Full-Stack JavaScript", "Why Learn Data Structures and Algorithms", "Best Coding Practices for Web Development", "How to Automate Testing in Your Projects", "A Complete Guide to Using GitHub Actions", 
		"How to Create an API with Python and Flask", "The Benefits of Using TypeScript", "Exploring the Use of WebAssembly", "How to Use WebRTC for Real-Time Communication", "Building a Chat App with React and Firebase", 
		"Why You Should Learn Go for Backend Development", "Introduction to RESTful API Authentication", "The Best Tools for Managing Dependencies in Your Projects", "Understanding REST and SOAP Web Services", 
		"How to Implement Pagination in Your API", "Why Python Is Great for Web Scraping", "How to Create a Responsive Web Design", "Best Practices for API Security", "Getting Started with Laravel Framework", 
		"How to Build a Portfolio with HTML and CSS", "Best Practices for Handling Errors in APIs", "Why Data Science Is Important for the Future", "How to Create a Blog with Flask", "Introduction to Docker Swarm", 
		"How to Secure Your Web Applications", "Building an Admin Dashboard with React", "How to Use Cloud Databases in Your Projects", "How to Build a Real-Time Web Application", "The Importance of Accessibility in Web Development", 
		"How to Build and Deploy a Python Web Application", "Exploring Cloud-Native Applications", "The Best Frontend Frameworks to Learn", "How to Work with APIs in Python", "Best Practices for Web Application Security", 
		"How to Build an API with Node.js and Express", "How to Use Event-Driven Architecture in Web Apps", "The Basics of Backend Development", "How to Create a RESTful Web Service", "The Complete Guide to Flask", 
		"Why is Continuous Delivery Important?", "Introduction to WebSockets in Web Development", "How to Create a Simple Web Application", "The Best Practices for Writing Unit Tests", "How to Use Redis for Caching", 
		"Why You Should Learn Web Development", "The Role of Data Science in Business", "How to Manage State in React", "The Importance of Code Quality", "Building Your First Server with Node.js", 
		"How to Optimize Your Application for SEO", "The Complete Guide to TypeScript", "Best Practices for Writing Code in Go", "How to Build a Simple API with Laravel", "Understanding the Basics of CSS", 
		"Getting Started with PHP", "How to Design RESTful APIs", "Best Practices for Designing User Interfaces", "How to Make Your Web Application Mobile-Friendly", "Exploring the Best Development Tools", 
		"How to Write a Secure Web Application", "The Future of Mobile Development", "Creating a Personal Blog with Django", "Best Practices for Building Scalable Applications", "The Complete Guide to Cloud Computing",
	}
	randomContents = []string{
		"Building web applications can be a daunting task, but with the right tools and practices, it becomes much more manageable. In this post, we will explore the essential steps to start building your own web app from scratch, including choosing the right framework, setting up a development environment, and deploying your app to the cloud.",
		"Clean code is essential for maintaining the long-term health of a project. Writing code that is easy to read, understand, and maintain is just as important as writing code that works. In this post, we will discuss some key practices for writing clean and efficient code that your team can easily maintain.",
		"Go programming is a powerful tool for building scalable, high-performance applications. This post will introduce you to the basics of Go and why it's a great language to learn for backend development. We'll also walk you through setting up a Go development environment and building your first Go application.",
		"Machine learning has become an essential tool for data analysis, prediction, and automation. In this beginner's guide, we'll cover the fundamentals of machine learning, including supervised and unsupervised learning, popular algorithms, and how to get started building your own machine learning models.",
		"SQL databases are a fundamental part of web development. In this post, we will explore the basics of SQL databases, including database design, normalization, and writing SQL queries to retrieve and manipulate data. By the end of this post, you will have a solid foundation in SQL.",
		"JavaScript is one of the most popular programming languages in the world. Whether you're building a web app or a server-side application, JavaScript can help you create dynamic and responsive user interfaces. In this post, we'll explore the basics of JavaScript and how to get started with writing JavaScript code.",
		"ReactJS is a powerful JavaScript library for building user interfaces. In this post, we'll explore some of the most useful ReactJS libraries you should know, including libraries for routing, state management, and animations. We'll also discuss how to integrate these libraries into your React app.",
		"Building RESTful APIs with Django is a great way to serve data to your frontend applications. In this post, we'll show you how to build a simple REST API using Django and Django REST Framework. We'll cover the basics of setting up a Django project, defining models, and creating API endpoints.",
		"SEO is crucial for driving traffic to your website. In this post, we'll explore the best practices for optimizing your website for search engines. From keyword research to technical SEO, we'll cover the key factors that can help you improve your website's visibility and ranking on search engines.",
		"GraphQL is a powerful query language for APIs. It allows you to fetch exactly the data you need, reducing the amount of data transferred between the client and the server. In this post, we'll explore the basics of GraphQL, how to set up a GraphQL API, and how to query data using GraphQL.",
		"Python is an incredibly versatile language that is great for data science, web development, automation, and more. In this post, we'll explore the key features of Python that make it such a great language to learn, and we'll provide you with some tips for getting started with Python programming.",
		"Cloud computing has transformed the way we build and deploy applications. In this post, we'll explore the benefits of cloud computing, including scalability, cost-efficiency, and flexibility. We'll also look at some of the most popular cloud platforms and how to get started with cloud development.",
		"Building a portfolio website is a great way to showcase your work and skills. In this post, we'll walk you through the process of creating a portfolio website from scratch, including how to design a simple yet elegant layout, how to add your projects, and how to host your site online.",
		"The future of artificial intelligence is incredibly exciting. AI has the potential to revolutionize nearly every industry, from healthcare to finance. In this post, we'll explore the latest developments in AI and how it is shaping the future of technology. We'll also discuss the ethical implications of AI and its impact on jobs.",
		"Blockchain technology has the potential to disrupt many industries, including finance, healthcare, and supply chain management. In this post, we'll explore the basics of blockchain, how it works, and the various use cases for blockchain technology in different industries.",
		"Software development can be challenging, but there are some common mistakes that can make the process even more difficult. In this post, we'll highlight five common mistakes that developers often make, and we'll provide tips for avoiding them. Whether you're a beginner or an experienced developer, this post will help you improve your workflow.",
		"Docker is a tool that allows you to easily deploy and manage applications in containers. In this post, we'll show you how to get started with Docker, including how to install Docker, create Dockerfiles, and manage Docker containers. We'll also discuss the benefits of using Docker for your projects.",
		"Git and GitHub are essential tools for version control and collaboration. In this post, we'll walk you through the basics of using Git and GitHub, including how to create a repository, clone a project, and make commits. We'll also discuss best practices for using Git and collaborating with other developers.",
		"Chatbots are becoming increasingly popular in customer service and automation. In this post, we'll show you how to build a simple chatbot using Python. We'll cover the basics of natural language processing (NLP), how to use a chatbot library, and how to integrate your chatbot into a website.",
		"Web accessibility is an essential aspect of web development that ensures your website can be used by people with disabilities. In this post, we'll explore the importance of web accessibility, including how to design for screen readers, how to make your website keyboard-friendly, and how to ensure your website meets accessibility standards.",
		"Kubernetes is a powerful tool for managing containerized applications in production. In this post, we'll explore the basics of Kubernetes, including how to set up a Kubernetes cluster, deploy containers, and scale applications. We'll also discuss how Kubernetes can help you manage complex, distributed applications.",
		"DevOps is a culture and set of practices that bring together development and operations teams to improve the delivery of software. In this post, we'll introduce you to the basics of DevOps, including continuous integration (CI), continuous deployment (CD), and infrastructure as code (IaC).",
		"Rust is a systems programming language that is known for its performance and memory safety. In this post, we'll explore why Rust is a great language to learn, how to get started with Rust programming, and some of the key features that make it stand out from other languages.",
		"Building an e-commerce website is a great way to learn about web development and online business. In this post, we'll walk you through the steps of building an e-commerce website, including how to design the user interface, set up a product catalog, and implement a shopping cart and checkout system.",
		"Open source software is software that is freely available for anyone to use, modify, and distribute. In this post, we'll explore the benefits of open source software, how to contribute to open source projects, and how you can get involved in the open source community.",
		"Serverless architecture is a cloud computing model where you can run your applications without managing servers. In this post, we'll explain the basics of serverless architecture, its benefits, and how to get started using serverless technologies like AWS Lambda and Azure Functions.",
		"React and Angular are two of the most popular frontend frameworks. In this post, we'll compare React and Angular, discussing their strengths and weaknesses, and help you decide which framework is right for your next project. Whether you're a beginner or an experienced developer, this post will give you the information you need to make an informed decision.",
		"Unit testing is an important part of software development that helps ensure your code works as expected. In this post, we'll explain the basics of unit testing, how to write unit tests in your preferred programming language, and the importance of test-driven development (TDD).",
		"Integrating payments into your website allows users to make purchases and donations online. In this post, we'll explore how to integrate a payment gateway like Stripe or PayPal into your website, including how to set up an account, implement payment forms, and handle transactions securely.",
		"Continuous integration (CI) is the practice of automatically testing and building your code every time a change is made. In this post, we'll explain the importance of CI, how to set up a CI pipeline, and the benefits of CI for your development workflow. We'll also discuss popular CI tools like Jenkins, GitLab CI, and CircleCI.",
		"Agile development is a project management methodology that focuses on iterative, incremental delivery of software. In this post, we'll introduce you to the basics of Agile development, including the Scrum framework, how to create user stories, and how to work in sprints to deliver software more efficiently.",
		"Data privacy is an important issue in today's digital world. In this post, we'll explore the importance of data privacy, including how to protect user data, comply with privacy regulations like GDPR, and implement security best practices to ensure your users' information is safe.",
		"Cybersecurity is essential for protecting your applications and data from malicious attacks. In this post, we'll discuss the basics of cybersecurity, including common threats like phishing and SQL injection, how to secure your website, and best practices for protecting sensitive data.",
		"Building an online blog with WordPress is an easy and effective way to share your thoughts with the world. In this post, we'll guide you through the process of setting up a blog with WordPress, including choosing a theme, creating posts, and customizing your website to fit your style.",
	}
	randomTags = []string{
		"web development", "clean code", "Go programming", "machine learning", "SQL databases", "JavaScript", "ReactJS", "Django", "SEO", "GraphQL", 
		"Python", "cloud computing", "portfolio website", "artificial intelligence", "blockchain", "software development", "Docker", "Git", "Backend",
		"chatbots", "web accessibility", "Kubernetes", "DevOps", "Rust", "e-commerce", "open source", "serverless architecture", "React", "Angular", 
		"unit testing", "payment integration", "continuous integration", "Agile development", "data privacy", "cybersecurity", "WordPress", "API", 
		"frontend development", "backend development", "web security", "data analysis", "cloud platforms", "containerization", "version control", 
		"CI/CD", "microservices", "web hosting", "UX/UI design", "testing frameworks", "cloud storage", "project management", "software design", 
		"technical SEO", "machine learning models", "server-side rendering", "scalable apps", "automation", "database management", "data science", 
		"programming best practices", "cloud infrastructure", "container orchestration", "server management", "continuous deployment", "data integrity", 
		"server-side apps", "React ecosystem", "modern web apps", "virtualization", "versioning", "CI pipelines", "JavaScript frameworks", "data models", 
		"app deployment", "website performance", "data manipulation", "web standards", "API design", "website optimization", "client-server architecture", 
		"CI tools", "team collaboration", "programming languages", "ReactJS components", "cloud security", "website analytics", "REST APIs", "responsive design", 
		"database queries", "responsive web design", "user experience", "content management systems", "cloud engineering", "app scalability", "website design", 
		"automation tools", "JavaScript libraries", "coding standards", "project collaboration", "continuous integration tools", "Docker containers", "Backend Engineering",
		"serverless computing", "web development tools", "tech trends", "DevOps tools", "software testing", "real-time applications", "cloud migration", 
		"programming patterns", "e-commerce development", "API testing", "data visualization", "cloud deployment", "mobile development", "UX research", 
		"open-source software", "network security", "software architecture", "cloud-based solutions", "machine learning algorithms", "cloud computing platforms",
	}
	randomComments = []string{
		"Great post! Really enjoyed reading it.",
		"Very insightful, I learned a lot from this.",
		"I disagree with some points, but overall it's a good read.",
		"Thanks for sharing this! It was exactly what I was looking for.",
		"Could you elaborate more on this topic? It's fascinating.",
		"Not bad, but I think you could add more examples.",
		"Really well written, keep up the good work!",
		"I found this post very helpful, thanks for the advice.",
		"This is a great topic. I'd love to see a part 2.",
		"Awesome! I will share this with my friends.",
		"I think you missed a few important points, though.",
		"Interesting perspective, but I have a different view on this.",
		"I've been searching for something like this. Thanks!",
		"Love this! Very easy to understand and follow.",
		"Could you clarify the last point a bit more? It's unclear.",
		"Great tips! I'll definitely try them out.",
		"I love how you structured the content. Very engaging.",
		"Totally agree with you. This is a great approach.",
		"Really enjoyed this post! I'll definitely check out the resources you mentioned.",
		"Solid post, though I wish you included more visuals.",
		"Very informative, I'll be bookmarking this.",
		"Great read, I'll be following your blog for more!",
		"I love the practical advice, it's very actionable.",
		"Very detailed post, thanks for taking the time to write this.",
		"I've been looking for similar articles, this one is by far the best.",
		"This is something I never considered before. Thanks for the new perspective.",
		"Good article but could use more examples for better clarity.",
		"I agree with most of what you said, but I think you missed a few things.",
		"Great content, but I think you should proofread a bit more.",
		"Nice article! It's got a lot of useful information.",
		"This was so helpful, I was struggling with this issue myself.",
		"Great job on breaking down the topic so clearly!",
		"I'll be referring to this post often. Thank you!",
		"Excellent post, would love to see more content on this.",
		"I really enjoyed your writing style. It was very easy to follow.",
		"Great content. Can't wait for the next post.",
		"Could you provide more case studies next time? They help a lot.",
		"This is a great start, but more depth would be appreciated.",
		"Wonderful post! Exactly what I was looking for.",
		"This was exactly what I needed, thank you for the clarity.",
		"Great read, thanks for the information. I learned a lot!",
		"I completely agree with your point on X, well said!",
		"This post has been bookmarked, great job!",
		"Awesome post, I shared it with my network.",
		"Would love to hear more about your personal experiences with this.",
		"I appreciate how in-depth you went into this topic.",
		"Very educational, I will definitely apply this in my work.",
		"Thanks for sharing this. It's incredibly useful.",
		"Great read, but I think there are a few things to reconsider.",
		"I think you could expand on this subject, it's fascinating.",
		"Nice explanation of a complicated topic, I'm impressed!",
		"Very thorough, I learned a lot more than I expected!",
		"Wonderful tips, I'll try these out for sure.",
		"This post is spot on. Really helpful, thank you.",
		"You've made some excellent points, I fully agree.",
		"This is a fantastic article. I've bookmarked it for future reference.",
		"I love how you explained this. It makes so much more sense now.",
		"I like the way you tackled this topic, very informative.",
		"Can you explain your approach in more detail? I'm interested.",
		"Great points! This post opened my eyes to new ideas.",
		"I have a similar experience, and this post really resonated with me.",
		"I think the examples could be more diverse, but it's a great post.",
		"Your post has motivated me to take action on this issue. Thanks!",
		"I found this post very useful and informative. Thanks for sharing.",
		"This post addresses all the issues I've been thinking about.",
		"Great job! This is a very comprehensive guide.",
		"Thanks for sharing your knowledge. This will help a lot of people.",
		"Your post gave me a lot to think about. Thanks for the insights!",
		"Very useful, I can see myself referring back to this.",
		"Excellent read. I appreciate the research you put into this post.",
		"Love the way you organized this content. Easy to follow.",
		"Very helpful, looking forward to more posts like this.",
		"This was an eye-opener, I didn't know this before!",
		"Great post, but I wish you included more statistics or data.",
		"Really good post! I think it would be better with more visuals though.",
		"Interesting article. I will be sharing this with my colleagues.",
		"Really informative! I'm glad I found this post.",
		"I'm so glad I found this post, it's exactly what I needed.",
		"This post is very thorough, I appreciate the time you put into it.",
		"Nice job! Would love to see more on this topic.",
		"Great read, thanks for sharing your thoughts on this.",
		"I'm really impressed with how well you explained this topic.",
		"Thank you for sharing, I've learned a lot from this article.",
		"Could you provide more examples next time? It would help a lot.",
		"I really liked the points you made here, it's very useful.",
		"I agree with your main points, but would love to see some alternatives.",
		"Nice work! This helped clear up a lot of confusion I had.",
		"I've bookmarked this post. It's exactly what I needed.",
		"Fantastic article! I love how you explained the concepts.",
		"Thanks for the great read! Can't wait to see more posts.",
		"Very informative, thanks for sharing these ideas.",
		"Great insights, I learned a lot. Thanks for sharing!",
		"I love how detailed this post is. Keep up the great work!",
		"Nice post, I would love to see some more posts like this.",
		"Excellent post! I'll be referring to this often.",
		"Very well written. It was easy to understand and very informative.",
		"Great tips! I can't wait to implement them in my workflow.",
		"This post is packed with useful information. Thanks!",
		"I really appreciate the research you've done for this post.",
		"Interesting perspective. I think there's a lot of value here.",
		"This is a great post! I'm sharing it with my friends.",
		"I loved this post! You've simplified a difficult topic.",
		"Fantastic post! I can't wait to read more from you.",
		"Such a helpful post, thank you for breaking things down.",
		"This post is extremely well-written, I've learned so much.",
		"Great job on this post! Keep it up.",
		"I really enjoyed this post! It's so helpful.",
		"Great post! I hope you write more content on this topic.",
		"Nice content, I look forward to reading more from you.",
		"Well done! This is a very informative post.",
		"I'm so glad I stumbled upon this post. Very insightful.",
		"This is a well-researched post. Thanks for sharing your findings.",
		"Very helpful, looking forward to more content like this.",
		"Great post, I found it extremely useful for my own work.",
		"Nice work on this article! I'll be reading more from you.",
		"This is exactly the kind of content I've been looking for!",
		"Very educational. Thanks for sharing all this knowledge.",
		"I think you've covered all the essential points in this post.",
		"Wonderful post, I'll definitely be following your blog.",
		"Really insightful, I can't wait to see what you post next.",
		"Great article. I think you explained everything really well.",
		"This is such an important topic. Glad you wrote about it.",
		"Your post gave me a lot of useful ideas. Thanks!",
		"Nice job explaining such a complex topic in simple terms.",
		"Great job! I'll be coming back to this post often.",
		"Really interesting post, I hope you write more on this topic.",
		"Great content! I think your tips will really help me.",
		"This was exactly what I was looking for. Thanks for the post!",
		"Great post, I'm looking forward to the next one!",
		"Such an informative post! Thanks for sharing your thoughts.",
		"Excellent job on this post. I learned so much.",
		"I enjoyed reading this! Very helpful and well-written.",
		"Well done! I think you provided a lot of useful information.",
		"Great work, I'll be following you for more insightful posts.",
		"I'm excited to apply what I've learned from this post!",
		"Very insightful, I learned so much from this article.",
		"Great job! I'm looking forward to reading more from you.",
		"Really helpful post, thanks for sharing your knowledge.",
		"Excellent post! I think I'll refer back to this often.",
		"This post helped me solve a problem I've been struggling with.",
		"Thanks for sharing your knowledge, it's very appreciated.",
		"Great post, very well written and informative.",
		"I've bookmarked this! It's one of the most useful articles I've read.",
		"Very interesting! I'll be checking out more posts from you.",
		"Nice post, I think you did a great job explaining everything.",
	}
)


// FUNCTIONS

// Seeding function
func Seed(Models models.Models) error {
	ctx := context.Background()

	// Generate users and insert them
	users := generateUsers(30)
	var createdUsers []*models.User
	for _, user := range users {
		// Creating user
		createdUser, err := Models.Users.Create(ctx, user)
		// Error handling
		if err != nil {
			fmt.Println("Error creating user: ", err)
			return err
		}
		createdUsers = append(createdUsers, createdUser) // store created user
	}
	fmt.Println("Users successfully seeded..")

	// Generate posts and insert them
	posts := generatePosts(50, createdUsers)
	var createdPosts []*models.Post
	for _, post := range posts {
		// Creating posts
		createdPost, err := Models.Posts.Create(ctx, post);
		// Error handling
		if err != nil {
			fmt.Println("Error creating post: ", err)
			return err
		}
		createdPosts = append(createdPosts, createdPost) // store created post
	}
	fmt.Println("Posts successfully seeded..")

	// Generate comments and insert them
	comments := generateComments(100, createdUsers, createdPosts)
	var createdComments []*models.Comment
	for _, comment := range comments {
		// Creating comments
		createdComment, err := Models.Comments.Create(ctx, comment)
		// Error handling
		if err != nil {
			fmt.Println("Error creating comment: ", err)
			return err
		}
		createdComments = append(createdComments, createdComment) // store created comment
	}
	fmt.Println("Comments successfully seeded..")

	// Generate replies and insert them
	replies := generateReplies(300, createdUsers, createdComments)
	for _, reply := range replies {
		// Creating replies
		_, err := Models.Comments.Create(ctx, reply)
		// Error handling
		if err != nil {
			fmt.Println("Error creating reply: ", err)
			return err
		}
	}
	fmt.Println("Replies successfully seeded..")

	return nil
}


// Generate n number of users
func generateUsers(num int) []*models.User {
	// Creating slice of default num number of User type
	users := make([]*models.User, num)

	// Looping through slice and updating users
	for i := 0; i < num; i++ {
		users[i] = &models.User{
			FirstName: randomNames[rand.Intn(len(randomNames))],
			LastName: randomNames[rand.Intn(len(randomNames))],
			Username: randomNames[i%len(randomNames)] + randomNumberString(),
			Email: randomNames[i%len(randomNames)] + randomNumberString() + "@gmail.com",
			Password: "12345",
		}
	}

	// return users slice
	return users
}


// Generate n number of posts
func generatePosts(num int, users []*models.User) []*models.Post {
	// Creating slice of default num number of Post type
	posts := make([]*models.Post, num)

	// creating num number of posts
	for i := 0; i < num; i++ {
		// Randomly picking author from users slice
		user := users[rand.Intn(len(users))]

		posts[i] = &models.Post{
			UserID: user.ID,
			Title: randomTitles[rand.Intn(len(randomTitles))],
			Content: randomContents[rand.Intn(len(randomContents))],
			Tags: generateRandomTags(),
		}
	}

	return posts
}


// Generate n number of comments
func generateComments(num int, users []*models.User, posts []*models.Post) []*models.Comment  {
	// Creating slice of default num number of Comment type
	comments := make([]*models.Comment, num)

	// creating num number of posts
	for i := 0; i < num; i++ {
		// Randomly picking user from users slice
		user := users[rand.Intn(len(users))]
		// Randomly picking post from posts slice
		post := posts[rand.Intn(len(posts))]

		comments[i] = &models.Comment{
			UserID: user.ID,
			PostID: sql.NullInt64{Int64: post.ID, Valid: true},
			ParentID: sql.NullInt64{Valid: false}, // No parent for top-level comments
			Content: randomComments[rand.Intn(len(randomComments))],
			Commenter: *user,
		}
	}

	return comments
}


// Generate n number of comments
func generateReplies(num int, users []*models.User, comments []*models.Comment) []*models.Comment  {
	// Creating slice of default num number of Comment type
	replies := make([]*models.Comment, num)
	
	for i := 0; i < num; i++ {
		// Randomly picking user from users slice
		user := users[rand.Intn(len(users))]
		// Randomly picking comment from comments slice
		comment := comments[rand.Intn(len(comments))]

		replies[i] = &models.Comment{
			UserID: user.ID,
			PostID: sql.NullInt64{Valid: false}, // No post for 2nd tier comments (replies to comments)
			ParentID: sql.NullInt64{Int64: comment.ID, Valid: true},
			Content: randomComments[rand.Intn(len(randomComments))],
			Commenter: *user,
		}
	}

	return replies
}


// Function to generate a random number of tags (max 5 tags)
func generateRandomTags() []string {
	// Generate a random number of tags (between 1 and 5)
	numTags := rand.Intn(6)

	// Create a slice to hold the selected tags
	selectedTags := make([]string, numTags)

	// Randomly select tags from randomTags
	for i := 0; i < numTags; i++ {
		selectedTags[i] = randomTags[rand.Intn(len(randomTags))]
	}

	return selectedTags
}


// Function to generateandom number string up to 6 digits
func randomNumberString() string {
	num := rand.Intn(1000000) // range: 0 to 999999
	return fmt.Sprintf("%06d", num) // pads with leading zeroes if needed
}




// Entry point

func main() {
	// Setting BASE path from this file
	BASE_PATH, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Setting env path
	ENV_PATH := filepath.Join(BASE_PATH, "internal/config/.env") // ENV file 
	
	// Setting server address
	DBURL, err := envloader.GetEnv(ENV_PATH, "DBURL")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Initializing new DB connection
	DB_CONNECTION, err := NewDBConnection(DBURL, 1, 1,"1m")

	if err != nil {
		fmt.Println("Error while initializing DB connection: ", err)
		return
	}

	// Your seeding logic here
	fmt.Println("Seeding the database...")
	Models := models.NewModels(DB_CONNECTION)
	Seed(Models);
}